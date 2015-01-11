package app

import (
    "net/http"
    "github/go-personal/views"
)

// route is a wrapper around http.HandleFunc that allows me to pass a string
// variable to my viewHandler functions
func route(path string, handler views.ViewHandler, routeName string) {
    http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
        handler(w, r, routeName)
   })
}

func staticRoute(staticPath string) {
    http.HandleFunc(staticPath, views.StaticFileHandler)
}

func init() {
    route("/", views.BasicPageView, "frontPage")
    route("/projects", views.BasicPageView, "projectPage")
    route("/essays", views.BasicPageView, "essayListPage")
    route("/about", views.BasicPageView, "underConstruction")
    route("/contact", views.BasicPageView, "underConstruction")
    staticRoute("/static/")
}
