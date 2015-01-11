package app

import (
    "net/http"
    "github/go-personal/views"
)

type viewHandler func(http.ResponseWriter, *http.Request, string)

// route is a wrapper around http.HandleFunc that allows me to pass a string
// variable to my viewHandler functions
func route(path string, handler viewHandler, refName string) {
    http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
        handler(w, r, refName)
   })
}

func staticRoute(staticPath string) {
    http.HandleFunc(staticPath, views.StaticFileHandler)
}