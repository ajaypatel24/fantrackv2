package handlers

import (
	"log"
	"net/http"
	"v3/config"
	"v3/services"

	"github.com/go-chi/render"
)

type TeamCompareHandler struct {
	TeamCompareService *services.TeamCompareService
}

func NewTeamCompareHandler(teamCompareService *services.TeamCompareService) *TeamCompareHandler {
	return &TeamCompareHandler{}
}

func (h *TeamCompareHandler) GetWinners(w http.ResponseWriter, r *http.Request) {
	store, err := config.Store.Get(r, "auth-session")

	token, _ := store.Values["token"]

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	data := h.TeamCompareService.GetData(token)

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, data)
	return
}

func (h *TeamCompareHandler) GetLeaders(w http.ResponseWriter, r *http.Request) {
	store, err := config.Store.Get(r, "auth-session")

	token, _ := store.Values["token"]

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	data := h.TeamCompareService.GetCategoryLeaders(token)

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, data)
	return
}
