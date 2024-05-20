package utils

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templates := template.Must(template.ParseGlob("templates/*.html"))
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
