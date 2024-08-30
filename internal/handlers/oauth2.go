package handlers

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"v3/config"
	"v3/services"
)

const oauthStateString = "state"

type OAuthHandler struct {
	OAuthService *services.OAuthService
}

func NewOAuthHandler(oAuthService *services.OAuthService) *OAuthHandler {
	return &OAuthHandler{OAuthService: oAuthService}
}

func (h *OAuthHandler) HandleYahooLogin(w http.ResponseWriter, r *http.Request) {
	// Redirect the user to Yahoo's OAuth2 login page
	//verifier := oauth2.GenerateVerifier()
	url := h.OAuthService.GetAuthUrl("state")
	log.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *OAuthHandler) HandleYahooCallback(w http.ResponseWriter, r *http.Request) {
	session, err := config.Store.Get(r, "auth-session")
	if err != nil {
		log.Print(err)
	}
	// Validate state
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Println("invalid oauth state, expected " + oauthStateString + ", got " + state + "\n")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Get the authorization code from the callback URL
	code := r.FormValue("code")
	token, err := h.OAuthService.ExchangeAuthCode(code)
	gob.Register(token)
	if err != nil {
		fmt.Printf("could not get token: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	log.Println("valid:", token.Valid())

	// Output the token to the console (or store it securely)

	// Use the token to make authenticated requests to Yahoo's API
	//gob.Register(oauth2.Token{}) //required to make session work
	session.Values["token"] = token
	err = session.Save(r, w)
	if err != nil {
		log.Println(err)
		return
	}

	//c, err := gc.GetAllTeamStats("428.l.27608", 1)
	http.Redirect(w, r, "/test", http.StatusFound)
	//fmt.Printf("%v", c)

}
