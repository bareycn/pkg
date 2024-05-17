package oauth2

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"net/http"
)

const (
	AuthURL  = "https://appleid.apple.com/auth/authorize"
	TokenURL = "https://appleid.apple.com/auth/token"
)

type AppleUser struct {
	Aud            string `json:"aud,omitempty"`
	AuthTime       int    `json:"auth_time,omitempty"`
	CHash          string `json:"c_hash,omitempty"`
	Exp            int    `json:"exp,omitempty"`
	Iat            int    `json:"iat,omitempty"`
	Iss            string `json:"iss,omitempty"`
	Nonce          string `json:"nonce,omitempty"`
	NonceSupported bool   `json:"nonce_supported,omitempty"`
	Sub            string `json:"sub,omitempty"`
	Email          string `json:"email,omitempty"`
	EmailVerified  bool   `json:"email_verified,omitempty"`
	IsPrivateEmail bool   `json:"is_private_email,omitempty"`
	RealUserStatus int    `json:"real_user_status,omitempty"`
}

type AppleProvider struct {
	Name           string
	config         *oauth2.Config
	authCodeOption []oauth2.AuthCodeOption
}

func NewApple(clientID, clientSecret, callbackUrl string, scopes ...string) Provider {
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  callbackUrl,
		Scopes:       []string{"name", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  AuthURL,
			TokenURL: TokenURL,
		},
	}

	if len(scopes) > 0 {
		config.Scopes = append(config.Scopes, scopes...)
	}

	authCodeOptions := make([]oauth2.AuthCodeOption, 0, 1)
	for _, scope := range config.Scopes {
		if scope == "email" || scope == "name" {
			authCodeOptions = append(authCodeOptions, oauth2.SetAuthURLParam("response_mode", "form_post"))
		}
	}

	return &AppleProvider{
		Name:           "apple",
		config:         config,
		authCodeOption: authCodeOptions,
	}
}

func (p *AppleProvider) AuthCodeURL(state string) string {
	url := p.config.AuthCodeURL(state, p.authCodeOption...)
	return url
}

func (p *AppleProvider) Exchange(code string) (*oauth2.Token, error) {
	return p.config.Exchange(context.Background(), code)
}

func (p *AppleProvider) UserInfo(token *oauth2.Token) (*User, error) {
	resp, err := http.Get(fmt.Sprintf("%s?access_token=%s", EndpointUserInfoUrl, token.AccessToken))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user info: %d", resp.StatusCode)
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var user AppleUser
	if err = json.Unmarshal(bytes, &user); err != nil {
		return nil, err
	}
	raw := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &raw); err != nil {
		return nil, err
	}
	return &User{
		//ID:            user.ID,
		//Name:          user.Name,
		//Provider:      "apple",
		//Email:         user.Email,
		//EmailVerified: user.EmailVerified,
		//FirstName:     user.GivenName,
		//LastName:      user.FamilyName,
		//NickName:      user.Name,
		//AvatarURL:     user.Picture,
		RawData: raw,
	}, nil
}

func (p *AppleProvider) RefreshToken(token *oauth2.Token) (*oauth2.Token, error) {
	return nil, nil
}
