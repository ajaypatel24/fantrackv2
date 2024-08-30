package config

import (
	"v3/api"

	"golang.org/x/oauth2"
)

var OAuthConfig *oauth2.Config

func InitAuthConfig() {

	OAuthConfig = api.GetOAuth2Config(
		"",
		"",
		"")

	// This random string is used to protect against CSRF attacks

}
