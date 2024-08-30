package main

import (
	"net/http"
	"v3/config"
	"v3/internal/router"
	"v3/services"
)

func main() {
	config.InitAuthConfig()
	config.InitSessionStore()

	oauthService := services.NewOAuthService()
	teamCompareService := services.NewTeamCompareService()
	r := router.RouterSetup(oauthService, teamCompareService)

	if err := http.ListenAndServe(":8080", r); err != nil {
	}

}
