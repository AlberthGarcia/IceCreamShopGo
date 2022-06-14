package main

import (
	"net/http"

	"github.com/AlberthGarcia/IceCreamShopGo/pkg/conf"
	"github.com/AlberthGarcia/IceCreamShopGo/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//routes has the routes of our web application and middlewares
func routes(app *conf.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	//routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
