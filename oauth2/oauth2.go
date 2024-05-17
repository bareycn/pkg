package oauth2

import (
	"errors"
	"golang.org/x/oauth2"
)

type Configuration struct {
	Name         string   `mapstructure:"name"`
	ClientID     string   `mapstructure:"client_id"`
	ClientSecret string   `mapstructure:"client_secret"`
	CallbackUrl  string   `mapstructure:"callback_url"`
	Scopes       []string `mapstructure:"scopes"`
}

type Provider interface {
	AuthCodeURL(state string) string
	Exchange(code string) (*oauth2.Token, error)
	UserInfo(token *oauth2.Token) (*User, error)
	RefreshToken(token *oauth2.Token) (*oauth2.Token, error)
}

type Providers map[string]Provider

var providers = make(Providers)

func New(conf ...Configuration) {
	for _, c := range conf {
		switch c.Name {
		case "google":
			providers[c.Name] = NewGoogle(c.ClientID, c.ClientSecret, c.CallbackUrl, c.Scopes...)
		case "apple":
			providers[c.Name] = NewApple(c.ClientID, c.ClientSecret, c.CallbackUrl, c.Scopes...)
		default:
			panic("provider not found")
		}
	}
}

func GetProvider(name string) (Provider, error) {
	provider, ok := providers[name]
	if !ok {
		return nil, errors.New("provider not found")
	}
	return provider, nil
}
