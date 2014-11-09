package controllers

import (
    "html/template"
    "net/http"
    "models"
)

// compile all templates and cache them
var templates = template.Must(template.ParseFiles(
    "templates/front.html",
    "templates/about.html",
    "templates/header.html",
    "templates/essay.html",
))

func RenderViewModelToTemplate(w http.ResponseWriter, templateName string, viewModel models.WebPage) {
    templates.ExecuteTemplate(w, templateName, viewModel)
}
