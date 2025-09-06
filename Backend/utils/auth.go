package auth

import (
	
	"time"

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