package router

import (
	"net/http"
	"v3/internal/handlers"
	custommiddle "v3/middleware"
	"v3/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RouterSetup(oauthService *services.OAuthService, teamCompareService *services.TeamCompareService) http.Handler {
	r := chi.NewRouter()
	oauthHandler := handlers.NewOAuthHandler(oauthService)
	teamCompareHandler := handlers.NewTeamCompareHandler(teamCompareService)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(custommiddle.AuthMiddleware(oauthHandler))
	r.Use(custommiddle.AuthMiddleware(teamCompareHandler))

	r.Get("/login", oauthHandler.HandleYahooLogin)
	r.Get("/callback", oauthHandler.HandleYahooCallback)

	r.Get("/test", teamCompareHandler.GetWinners)

	r.Get("/winners", teamCompareHandler.GetLeaders)

	//.Get("/de", teamCompareHandler.Te)

	return r

}
