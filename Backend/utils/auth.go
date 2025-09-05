package auth

import (
"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)


const( AccessTokenValidityPeriod  time.Duration = 2*time.Hour
RefreshTokenValidityPeriod int            =7*24*60*60
  RefreshTokenValidityPeriod          time.Duration = time.Duration(RefreshTokenValidityPeriodInSeconds) * time.Second
)

type Tokens struct (
	AccessToken string jsson:"access_token"
	RefreshToken string `json:"refresh_token"`
)

type Claims struct (
	*jwt.StandardClaims
	Custom interface{} `json:"custom"`

)

func NewAuthClient(accessTokenSecret string, refreshTokenSecret string) *AuthClient {
	return &AuthClient{
		accessTokenSecret:  accessTokenSecret,
		refreshTokenSecret: refreshTokenSecret,
	}
}