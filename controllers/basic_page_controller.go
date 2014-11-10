package controllers

import (
    "net/http"
    "models"
)

var templateTitleMap = map[string]string{
    "frontPage" : "Home",
    "aboutPage" : "About Me",
}

func RenderBasicPage(w http.ResponseWriter, templateName string) {
    basicPageModel := models.WebPage{
        Title : templateTitleMap[templateName],
    }
    templates.ExecuteTemplate(w, templateName, basicPageModel)
}