package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	AccessTokenValidityPeriod           time.Duration = 2 * time.Hour
	RefreshTokenValidityPeriodInSeconds int           = 7 * 24 * 60 * 60
	RefreshTokenValidityPeriod          time.Duration = time.Duration(RefreshTokenValidityPeriodInSeconds) * time.Second
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Client struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
}

func NewAuthClient(
	accessTokenSecret string,
	refreshTokenSecret string,
) *Client {
	return &Client{
		AccessTokenSecret:  accessTokenSecret,
		RefreshTokenSecret: refreshTokenSecret,
	}
}

// Claims combines RegisteredClaims with custom claims
type Claims struct {
	jwt.RegisteredClaims
	Custom any `json:"custom,omitempty"`
}

// createToken generates a signed JWT token
func createToken(signKey string, customClaims any, validityPeriod time.Duration) (string, error) {
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(validityPeriod)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "erp",
			Subject:   "auth",
		},
		Custom: customClaims,
	}

	// Use HS256 for signing
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign with the secret key
	return token.SignedString([]byte(signKey))
}
