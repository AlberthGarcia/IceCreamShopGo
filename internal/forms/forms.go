package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}

//New create a new empty error
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//HasFields verify if exists the field in the form
func (f Form) HasFields(fields string, r *http.Request) bool {
	existsField := r.Form.Get(fields)
	if existsField == "" {

		return false
	}

	return true
}
