package oauth2

import (
	"context"
	"fmt"
	"github.com/go-jose/go-jose/v4/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"net/http"
)

type GoogleUser struct {
	ID            string `json:"id,omitempty"`
	Iss           string `json:"iss,omitempty"`
	Azp           string `json:"azp,omitempty"`
	Aud           string `json:"aud,omitempty"`
	Sub           string `json:"sub,omitempty"`
	Email         string `json:"email,omitempty"`
	EmailVerified bool   `json:"email_verified,omitempty"`
	AtHash        string `json:"at_hash,omitempty"`
	Name          string `json:"name,omitempty"`
	Picture       string `json:"picture,omitempty"`
	GivenName     string `json:"given_name,omitempty"`
	FamilyName    string `json:"family_name,omitempty"`
	Locale        string `json:"locale,omitempty"`
	Iat           int    `json:"iat,omitempty"`
	Exp           int    `json:"exp,omitempty"`
}

type GoogleProvider struct {
	Name           string
	config         *oauth2.Config
	authCodeOption []oauth2.AuthCodeOption
}

const EndpointUserInfoUrl = "https://www.googleapis.com/oauth2/v2/userinfo"

func NewGoogle(clientID, clientSecret, callbackUrl string, scopes ...string) Provider {
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  callbackUrl,
		Scopes:       []string{"profile", "email", "openid"},
		Endpoint:     google.Endpoint,
	}

	if len(scopes) > 0 {
		config.Scopes = append(config.Scopes, scopes...)
	}

	return &GoogleProvider{
		Name:           "google",
		config:         config,
		authCodeOption: []oauth2.AuthCodeOption{oauth2.AccessTypeOffline},
	}
}

func (g *GoogleProvider) AuthCodeURL(state string) string {
	return g.config.AuthCodeURL(state, g.authCodeOption...)
}

func (g *GoogleProvider) Exchange(code string) (*oauth2.Token, error) {
	return g.config.Exchange(context.Background(), code)
}

func (g *GoogleProvider) UserInfo(token *oauth2.Token) (*User, error) {
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
	var user GoogleUser
	if err = json.Unmarshal(bytes, &user); err != nil {
		return nil, err
	}
	raw := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &raw); err != nil {
		return nil, err
	}
	return &User{
		ID:            user.ID,
		Name:          user.Name,
		Provider:      "google",
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		FirstName:     user.GivenName,
		LastName:      user.FamilyName,
		NickName:      user.Name,
		AvatarURL:     user.Picture,
		RawData:       raw,
	}, nil
}

func (g *GoogleProvider) RefreshToken(token *oauth2.Token) (*oauth2.Token, error) {
	panic("implement me")
}
