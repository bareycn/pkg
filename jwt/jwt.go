package jwt

import (
	"github.com/bareycn/pkg/r"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"strconv"
	"time"
)

const ContextUserID = "userID"

var conf *Configuration

type Configuration struct {
	Secret string        `mapstructure:"secret"`
	Expire time.Duration `mapstructure:"expire"`
	Issuer string        `mapstructure:"issuer"`
}

type TokenResponse struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	TokenType    string        `json:"token_type"`
	ExpireIn     time.Duration `json:"expire_in"`
}

func New(config Configuration) {
	conf = &config
}

func GenerateAccessToken(userID int64) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   strconv.FormatInt(userID, 10),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(conf.Expire * time.Second)),
		Issuer:    conf.Issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.Secret))
}

// Auth jwt中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			return []byte(conf.Secret), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(401, r.Error(err))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			c.AbortWithStatusJSON(401, r.Error(err))
			return
		}
		userID, err := claims.GetSubject()
		if err != nil {
			c.AbortWithStatusJSON(401, r.Error(err))
			return
		}
		c.Set(ContextUserID, userID)
		c.Next()
	}
}
