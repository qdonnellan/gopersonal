package routers

import (
    "net/http"
    "views"
)

func init() {
    http.HandleFunc("/", views.FrontPageView)
    http.HandleFunc("/about", views.AboutPageView)
    http.HandleFunc("/essays/", views.IndividualEssayPageView)
    http.HandleFunc("/static/", views.StaticFileViewHandler)
}
