package views

import (
    "net/http"
    "html/template"
)

// compile all templates and cache them
var templates = template.Must(template.ParseFiles(
    "templates/front.html",
    "templates/about.html",
    "templates/header.html",
))

// Handle the front page view of the website
func FrontPageView(w http.ResponseWriter, r *http.Request) {
    templates.ExecuteTemplate(w, "frontPage", frontPageViewModel)
}

// Handle the about page view of the website
func AboutPageView(w http.ResponseWriter, r *http.Request) {
    templates.ExecuteTemplate(w, "aboutPage", aboutPageViewModel)
}

// Handle all static files according to their name
func StaticFileViewHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
}



