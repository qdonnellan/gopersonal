package router

import (
    "net/http"
    "gopersonal/essays"
    "gopersonal/website"
)

func init() {
    http.HandleFunc("/", website.FrontPageHandler)
    http.HandleFunc("/essays", essays.Handler)
}
