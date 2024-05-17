package oidc

import (
	"github.com/bareycn/pkg/oauth2"
	"github.com/coreos/go-oidc/v3/oidc"
)

// ParseGoogleClaims 解析Google Claims
func ParseGoogleClaims(idToken *oidc.IDToken) (*oauth2.User, error) {
	var user oauth2.GoogleUser
	if err := idToken.Claims(&user); err != nil {
		return nil, err
	}

	var raw map[string]interface{}
	if err := idToken.Claims(&raw); err != nil {
		return nil, err
	}

	return &oauth2.User{
		ID:            user.Sub,
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
