package utils

import (
	"html/template"
	"net/http"
)


var templates *template.Template

//CarregarTemplates coloca os templates html na variavel templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
	
}
