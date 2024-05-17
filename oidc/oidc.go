package oidc

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/bareycn/pkg/oauth2"
	"github.com/coreos/go-oidc/v3/oidc"
	"log"
)

// Provider OIDC认证提供者
type Provider struct {
	Name                 string
	verify               *oidc.IDTokenVerifier
	SkipAccessTokenCheck bool
	SkipNonceCheck       bool
}

const (
	IssuerGoogle = "https://accounts.google.com"
	IssuerApple  = "https://appleid.apple.com"
)

type Providers map[string]Provider

var providers = make(Providers)

// New 创建OIDC认证提供者
func New() {
	log.Println("等待初始化OIDC认证提供者")
	// 注册OIDC服务
	googleOidc, err := NewProvider("google", IssuerGoogle, false, true)
	if err != nil {
		panic(err)
	}
	appleOidc, err := NewProvider("apple", IssuerApple, true, false)
	if err != nil {
		panic(err)
	}
	UseProvider(googleOidc, appleOidc)
	log.Println("OIDC认证提供者初始化完成")
}

// NewProvider 创建OIDC认证提供者
func NewProvider(name, issuer string, skipAccessTokenCheck bool, skipNonceCheck bool) (Provider, error) {
	provider, err := oidc.NewProvider(context.Background(), issuer)
	if err != nil {
		return Provider{}, err
	}
	verifier := provider.Verifier(&oidc.Config{
		SkipClientIDCheck: true,
	})

	return Provider{
		Name:                 name,
		verify:               verifier,
		SkipAccessTokenCheck: skipAccessTokenCheck,
		SkipNonceCheck:       skipNonceCheck,
	}, nil
}

// UseProvider 注册OIDC认证提供者
func UseProvider(provider ...Provider) {
	for _, p := range provider {
		providers[p.Name] = p
	}
}

// GetProvider 获取OIDC认证提供者
func GetProvider(name string) (Provider, error) {
	provider := providers[name]
	if provider.Name == "" {
		return Provider{}, errors.New("provider not found")
	}
	return provider, nil
}

// VerifyToken 验证ID Token
func (p *Provider) VerifyToken(token string, accessToken string, nonce string) (*oidc.IDToken, error) {
	idToken, err := p.verify.Verify(context.Background(), token)
	if err != nil {
		return nil, err
	}
	if !p.SkipAccessTokenCheck {
		if err = idToken.VerifyAccessToken(accessToken); err != nil {
			return nil, err
		}
	}
	if !p.SkipNonceCheck {
		nonceHash := idToken.Nonce != ""
		nonceParam := nonce != ""
		if nonceHash != nonceParam {
			return nil, errors.New("nonce not found")
		} else if nonceHash && nonceParam {
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(nonce)))
			if idToken.Nonce != hash {
				return nil, errors.New("nonce not match")
			}
		}
	}
	return idToken, nil
}

// GetUserInfo  获取用户信息
func (p *Provider) GetUserInfo(idToken *oidc.IDToken) (*oauth2.User, error) {
	switch p.Name {
	case "google":
		return ParseGoogleClaims(idToken)
	case "apple":
		return ParseAppleClaims(idToken)
	}
	return nil, errors.New("provider not found")
}
