package controllers

import (
    "html/template"
    "net/http"
    "github/go-personal/models"
)

// compile all templates and cache them
var templates = template.Must(template.ParseGlob("templates/*"))


// RenderTeamplateFromPageModel finds a template with the specified
// TemplateName and renders that template from data in pageModel.
func RenderTemplateFromPageModel(w http.ResponseWriter, pageModel models.WebPage) {
    templates.ExecuteTemplate(w, pageModel.TemplateName, pageModel)
}