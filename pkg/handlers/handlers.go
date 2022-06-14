package handlers

import (
	"net/http"

	"github.com/AlberthGarcia/IceCreamShopGo/pkg/conf"
	"github.com/AlberthGarcia/IceCreamShopGo/pkg/models"
	"github.com/AlberthGarcia/IceCreamShopGo/pkg/render"
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
	render.RenderTemplates(w, "index.page.html", &models.TemplateData{})
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	strignMapData := make(map[string]string)
	strignMapData["Alberth"] = "Test"

	//get the ip address and save in the map
	remoteIP := repo.AppConfig.Session.GetString(r.Context(), "remote_ip")
	strignMapData["remote_ip"] = remoteIP

	render.RenderTemplates(w, "about.page.html", &models.TemplateData{
		StringMap: strignMapData,
	})
}
