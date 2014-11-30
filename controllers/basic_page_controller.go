package controllers

import (
    "net/http"
    "github/go-personal/models"
)

var templateTitleMap = map[string]string{
    "frontPage" : "Quentin Donnellan: Front Page",
    "aboutPage" : "About Quentin Donnellan",
    "essayErrorPage" : "Error Loading Essay",
    "constructionPage" : "Under Construction",
    "essayListPage" : "All Essays",
    "projectPage" : "Side Projects created by Quentin Donnellan",
}

func RenderBasicPage(w http.ResponseWriter, templateName string) {
    basicPageModel := models.WebPage{
        Title : templateTitleMap[templateName],
        TemplateName : templateName,
    }
    RenderTemplateFromPageModel(w, basicPageModel)
}