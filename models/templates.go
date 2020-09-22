package models

import (
	"fmt"
	"html/template"
	"net/http"
)

func LoadTemplates(pattern string) (templates *template.Template) {
	templates = template.Must(template.ParseGlob(pattern))
	return templates
}

func ExecuteTemplate(w http.ResponseWriter, templates *template.Template, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		fmt.Println("Error in template execution --> ", err.Error())
	}
	return
}
