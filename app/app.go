package app

import (
    "net/http"
    "github/go-personal/views"
)

func init() {
    http.HandleFunc("/", views.FrontPageView)
    http.HandleFunc("/about", views.ConstructionPageView)
    http.HandleFunc("/contact", views.ConstructionPageView)
    http.HandleFunc("/projects", views.ProjectPageView)
    http.HandleFunc("/essays/", views.EssayPageView)
    http.HandleFunc("/static/", views.StaticFileViewHandler)
}
