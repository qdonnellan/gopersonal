package views

import (
    "net/http"
    "controllers"
)

// Handle the front page
func FrontPageView(w http.ResponseWriter, r *http.Request) {
    controllers.RenderBasicPage(w, "frontPage")
}

// Handle the about page
func AboutPageView(w http.ResponseWriter, r *http.Request) {
    controllers.RenderBasicPage(w, "aboutPage")
}

// Handle an individual essay page
func IndividualEssayPageView(w http.ResponseWriter, r *http.Request) {
    // skip the first characters: /essays/ (length 8)
    essayTitle := r.URL.Path[8:]
    controllers.RenderEssayPage(w, essayTitle)
}

// Handle all static files according to their name
func StaticFileViewHandler(w http.ResponseWriter, r *http.Request) {
    staticFileName := r.URL.Path[1:]
    http.ServeFile(w, r, staticFileName)
}



