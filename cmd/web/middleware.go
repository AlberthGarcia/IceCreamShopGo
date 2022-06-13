package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//NoSurf build a csrfHandler and set a cookie to make more secure the requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
