package main

import (
	"net/http"

	"github.com/AlberthGarcia/IceCreamShopGo/pkg/handlers"
)

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.ListenAndServe(":8080", nil)
}
