package main

import (
	"log"
	"net/http"

	"github.com/AlberthGarcia/IceCreamShopGo/pkg/conf"
	"github.com/AlberthGarcia/IceCreamShopGo/pkg/handlers"
	"github.com/AlberthGarcia/IceCreamShopGo/pkg/render"
)

const numberPort = ":8080"

func main() {
	var appConfig conf.AppConfig

	//call our function CreateTemplateCache
	templateCache, err := render.CreateTemplateMapCache()
	if err != nil {
		log.Fatal("Cannot create the template cache", err)
	}
	//save in our app struct our templates cache
	appConfig.TemplateCache = templateCache
	appConfig.UseCache = false

	//set the appConfig full and get a Repo
	repo := handlers.SetRepository(&appConfig)
	//set the repo to save it in the struct
	handlers.NewHandler(repo)
	render.NewTemplate(&appConfig)

	server := &http.Server{
		Addr:    numberPort,
		Handler: routes(&appConfig),
	}

	err = server.ListenAndServe()
	log.Fatal(err)

}
