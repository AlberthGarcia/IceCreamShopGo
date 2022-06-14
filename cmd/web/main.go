package main

import (
	"log"
	"net/http"
	"time"

	"github.com/AlberthGarcia/IceCreamShopGo/pkg/conf"
	"github.com/AlberthGarcia/IceCreamShopGo/pkg/handlers"
	"github.com/AlberthGarcia/IceCreamShopGo/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const numberPort = ":8080"

var appConfig conf.AppConfig
var session *scs.SessionManager

func main() {
	//change this to true when in production
	appConfig.InProduction = false

	//initialize sessions
	session = scs.New()
	//Configure my cookie session
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.InProduction

	//save our session in our config struct
	appConfig.Session = session

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
