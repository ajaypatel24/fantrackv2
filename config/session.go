package config

import "github.com/gorilla/sessions"

var Store *sessions.CookieStore

func InitSessionStore() {
	Store = sessions.NewCookieStore([]byte("secret-key"))
}
