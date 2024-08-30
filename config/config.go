package config

import (
	"log"
	"os"
	"v3/api"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

var OAuthConfig *oauth2.Config

func InitAuthConfig() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Failed to load ENV")
	}

	OAuthConfig = api.GetOAuth2Config(
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
		os.Getenv("REDIRECT_URL"))

}
