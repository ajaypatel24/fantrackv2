package services

import (
	"context"
	"golang.org/x/oauth2"
	"v3/api"
	"v3/config"
)

type OAuthService struct{}

func NewOAuthService() *OAuthService {
	return &OAuthService{}
}

func (s *OAuthService) GetAuthUrl(state string) string {
	return config.OAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (s *OAuthService) ExchangeAuthCode(code string) (*oauth2.Token, error) {

	return config.OAuthConfig.Exchange(context.Background(), code)
}

func (s *OAuthService) GetClient(token any) *api.Client {
	oauthClient := config.OAuthConfig.Client(context.Background(), token.(*oauth2.Token))
	return api.NewClient(oauthClient)
}
