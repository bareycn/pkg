package oidc

import (
	"github.com/bareycn/pkg/oauth2"
	"github.com/coreos/go-oidc/v3/oidc"
)

// ParseAppleClaims 解析Apple Claims
func ParseAppleClaims(idToken *oidc.IDToken) (*oauth2.User, error) {
	var user oauth2.AppleUser
	if err := idToken.Claims(&user); err != nil {
		return nil, err
	}

	var raw map[string]interface{}
	if err := idToken.Claims(&raw); err != nil {
		return nil, err
	}

	return &oauth2.User{
		ID:            user.Sub,
		Provider:      "apple",
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		RawData:       raw,
	}, nil
}
