package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

//RenderTemplates render templates using html/templates
func RenderTemplates(w http.ResponseWriter, tmp string) {
	//Call our function to render the templates
	templateCache, err := CreateTemplateMapCache()
	if err != nil {
		log.Println("Error getting the templates")
		return
	}

	//verify that the template exists
	temp, ok := templateCache[tmp]
	if !ok {
		log.Fatal("Template does not exists", err)
		return
	}

	bytesTemp := new(bytes.Buffer)

	//execute the template and save its value into bytesTemp with any Data
	err = temp.Execute(bytesTemp, nil)
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
