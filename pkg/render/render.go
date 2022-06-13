package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/AlberthGarcia/IceCreamShopGo/pkg/conf"
	"github.com/AlberthGarcia/IceCreamShopGo/pkg/models"
)

var functions = template.FuncMap{}
var appConfig *conf.AppConfig

//NewTemplate sets the config to use it in this file
func NewTemplate(appConf *conf.AppConfig) {
	appConfig = appConf
}

//AddDefaultData
func AddDefaultData(data *models.TemplateData) *models.TemplateData {

	return data
}

//RenderTemplates render templates using html/templates
func RenderTemplates(w http.ResponseWriter, tmp string, data *models.TemplateData) {

	var templateCache map[string]*template.Template
	//True-> we're going to use the app struct to load the templates
	if appConfig.UseCache {
		templateCache = appConfig.TemplateCache
	} else {
		//read each template since the disk
		templateCache, _ = CreateTemplateMapCache()
	}

	//verify that the template exists
	temp, ok := templateCache[tmp]
	if !ok {
		log.Fatal("Template does not exists")
		return
	}

	//new Var type of bytes.Buffer
	bytesTemp := new(bytes.Buffer)

	//func to add default data
	data = AddDefaultData(data)
	//execute the template and save its value into bytesTemp with any Data
	err := temp.Execute(bytesTemp, data)
	if err != nil {
		log.Println("Cannot execute the template", err)
		return
	}

	//write the bytes in the response writer
	_, err = bytesTemp.WriteTo(w)
	if err != nil {
		log.Fatal("Error writing the template to browser", err)
		return
	}

}

//CreateTemplateMapCache create a map with all templates in our folder templates
func CreateTemplateMapCache() (map[string]*template.Template, error) {
	mapTemplateCache := map[string]*template.Template{}

	//get all files with this pattern
	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return mapTemplateCache, err
	}

	//range page per page
	for _, page := range pages {

		//get the path name of the page
		namePage := filepath.Base(page)
		//create our template of this pages
		ts, err := template.New(namePage).Funcs(functions).ParseFiles(page)
		if err != nil {
			return mapTemplateCache, err
		}

		//verify if in this path there are files with this pattern
		matchPage, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return mapTemplateCache, err
		}

		//if there is more than 1
		if len(matchPage) > 0 {
			//save in our templates (TS) the file that match with this pattern
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return mapTemplateCache, err
			}
		}
		//save in our map of templates their names with their reference
		mapTemplateCache[namePage] = ts
	}

	return mapTemplateCache, nil

}
