package controllers

import (
    "net/http"
    "github/go-personal/models"
    "io/ioutil"
)

// RenderEssayPage renders the essay page template by populating
// it with the associated markdown content.
func RenderEssayPage(w http.ResponseWriter, essayTitle string) {
    essayContent, err := getEssayContentByEssayTitle(essayTitle)
    if err != nil {
        RenderBasicPage(w, "essayErrorPage")
    } else {
        essayPageModel := createEssayPageModel(essayTitle, essayContent)
        RenderTemplateFromPageModel(w, essayPageModel)
    }
}

func getEssayContentByEssayTitle(essayTitle string) (string, error) {
    filePath := "content/essays/" + essayTitle + ".md"
    essayContent, err := ioutil.ReadFile(filePath)
    return string(essayContent), err
}

func createEssayPageModel(title string, content string) models.WebPage {
    return models.WebPage{
        Title : title,
        MainContent : content,
        TemplateName : "essayPage",
    }
}
