package app

import "github/go-personal/views"

func init() {
    route("/", views.BasicPageView, "frontPage")
    route("/projects", views.BasicPageView, "projectPage")
    route("/essays", views.BasicPageView, "essayListPage")
    route("/about", views.BasicPageView, "underConstruction")
    route("/contact", views.BasicPageView, "underConstruction")
    route("/essays/", views.EssayPageView, "essayPage")
    staticRoute("/static/")
}
