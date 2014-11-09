package views

import (
    "net/http"
    "html/template"
    "controllers"
)

// compile all templates and cache them
var templates = template.Must(template.ParseFiles(
    "templates/front.html",
    "templates/about.html",
    "templates/header.html",
))

// Handle the front page
func FrontPageView(w http.ResponseWriter, r *http.Request) {
    templates.ExecuteTemplate(w, "frontPage", frontPageViewModel)
}

// Handle the about page
func AboutPageView(w http.ResponseWriter, r *http.Request) {
    templates.ExecuteTemplate(w, "aboutPage", aboutPageViewModel)
}

// Handle an individual essay page
func IndividualEssayPageView(w http.ResponseWriter, r *http.Request) {
    // skip the first characters: /essays/ (length 8)
    essayTitle := r.URL.Path[8:]
    controllers.RenderEssayTemplate(w, essayTitle)
}

// Handle all static files according to their name
func StaticFileViewHandler(w http.ResponseWriter, r *http.Request) {
    staticFileName := r.URL.Path[1:]
    http.ServeFile(w, r, staticFileName)
}



