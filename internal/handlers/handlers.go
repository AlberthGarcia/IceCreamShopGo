package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AlberthGarcia/IceCreamShopGo/internal/conf"
	"github.com/AlberthGarcia/IceCreamShopGo/internal/models"
	"github.com/AlberthGarcia/IceCreamShopGo/internal/render"
)

//Repo is a var of Repository
var Repo *Repository

//Repository handlers our controllers
type Repository struct {
	AppConfig *conf.AppConfig
}

//Set the repository to the struct and return it full
func SetRepository(app *conf.AppConfig) *Repository {
	return &Repository{
		AppConfig: app,
	}
}

//NewHandler set the var Repo, get in the argument to the repo var
func NewHandler(rep *Repository) {
	Repo = rep
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//get the ip of the user
	remoteIP := r.RemoteAddr
	//pull that ip address in our session
	repo.AppConfig.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplates(w, r, "index.page.html", &models.TemplateData{})
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	strignMapData := make(map[string]string)
	strignMapData["Alberth"] = "Test"

	//get the ip address and save in the map
	remoteIP := repo.AppConfig.Session.GetString(r.Context(), "remote_ip")
	strignMapData["remote_ip"] = remoteIP

	render.RenderTemplates(w, r, "about.page.html", &models.TemplateData{
		StringMap: strignMapData,
	})
}

//Contact render the page Contact
func (repo *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, r, "contact.page.html", &models.TemplateData{})
}

//PostContact get the fields in the form
func (repo *Repository) PostContact(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		us := r.Form.Get("username")
		ps := r.Form.Get("password")
		to := r.Form.Get("csrf_token")

		w.Write([]byte(fmt.Sprintf("Hola %s your password is: %s and token is: %s", us, ps, to)))
	}

}

//jsonResponse struct to response with json
type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

//JsonContact send the conf to the header and return a json with the response
func (repo *Repository) JsonContact(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Avilable",
	}

	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
