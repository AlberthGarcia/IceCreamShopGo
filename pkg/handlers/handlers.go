package handlers

import (
	"net/http"

	"github.com/AlberthGarcia/IceCreamShopGo/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "index.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplates(w, "about.page.html")
}
