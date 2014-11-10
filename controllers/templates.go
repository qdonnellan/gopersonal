package controllers

import "html/template"

// compile all templates and cache them
var templates = template.Must(template.ParseFiles(
    "templates/front.html",
    "templates/about.html",
    "templates/header.html",
    "templates/essay.html",
))