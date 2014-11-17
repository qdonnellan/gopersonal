package controllers

import (
    "net/http"
    "models"
    "io/ioutil"
)

func RenderEssayPage(w http.ResponseWriter, essayTitle string) {
    markdownContent, err := getEssayContentByEssayTitle(essayTitle)
    if err != nil {
        RenderBasicPage(w, "essayErrorPage")
    } else {
        essayPageModel := createEssayPageModel(essayTitle, markdownContent)
        templates.ExecuteTemplate(w, "essayPage", essayPageModel)
    }
}

func getEssayContentByEssayTitle(essayTitle string) (string, error) {
    filePath := "content/essays/" + essayTitle + ".md"
    markdownData, err := ioutil.ReadFile(filePath)
    return string(markdownData), err
}

func createEssayPageModel(title string, content string) models.WebPage {
    return models.WebPage{
        Title : title,
        MainContent : content,
    }
}
