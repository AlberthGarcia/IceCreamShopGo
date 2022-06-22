package main

import (
	"net/http"

	"github.com/AlberthGarcia/IceCreamShopGo/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//routes has the routes of our web application and middlewares
func routes() http.Handler {
	mux := chi.NewRouter()

	//middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	//routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	//url-contact
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Post("/contact", handlers.Repo.PostContact)
	mux.Post("/contact-json", handlers.Repo.JsonContact)

	return mux
}
