package views

import (
    "net/http"
    "github/go-personal/controllers"
)

// ViewHandler describes what each view handler function should look like
// such that custom wrapping of http.HandleFunc is predictable (see app.go)
type ViewHandler func(http.ResponseWriter, *http.Request, string)

func BasicPageView(w http.ResponseWriter, r *http.Request, pageName string) {
    controllers.RenderBasicPage(w, pageName) 
}

// EssayPageView calls the RenderEssayPage controller to render an individual
// essay.
func EssayPageView(w http.ResponseWriter, r *http.Request) {
    // skip the first characters: /essays/ (length 8)
    essayTitle := r.URL.Path[8:]
    controllers.RenderEssayPage(w, essayTitle)
}

// TutorialPageView calls the RenderEssayPage controller to render an individual
// tutorial.
func TutorialPageView(w http.ResponseWriter, r *http.Request) {
    // skip the first characters: /essays/ (length 8)
    essayTitle := r.URL.Path[8:]
    controllers.RenderEssayPage(w, essayTitle)
}

// StaticFileHandler handles all static file requests. 
func StaticFileHandler(w http.ResponseWriter, r *http.Request) {
    staticFileName := r.URL.Path[1:]
    http.ServeFile(w, r, staticFileName)
}