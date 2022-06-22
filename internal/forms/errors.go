package forms

type errors map[string][]string

//AddError add an error for each form field
func (e errors) AddError(field, message string) {
	e[field] = append(e[field], message)
}

//GetError gets an error through his key
func (e errors) GetError(field string) string {
	errorString := e[field]
	if len(errorString) == 0 {
		return ""
	}
	return errorString[0]

}
