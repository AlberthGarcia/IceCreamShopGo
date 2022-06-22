package models

//TemplateData has the fields to use to send to the template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Success   string
	Error     string
	Warning   string
}
