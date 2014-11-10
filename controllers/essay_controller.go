package controllers

import (
    "net/http"
    "models"
    "io/ioutil"
)

func RenderEssayPage(w http.ResponseWriter, essayTitle string) {
    markdownContent := getEssayContentByEssayTitle(essayTitle)
    essayPageModel := createEssayPageModel(essayTitle, markdownContent)
    templates.ExecuteTemplate(w, "essayPage", essayPageModel)
}

func getEssayContentByEssayTitle(essayTitle string) string {
    filePath := "content/essays/" + essayTitle + ".md"
    markdownData, err := ioutil.ReadFile(filePath)
    if err != nil {
        panic(err)
    } else {
        return string(markdownData)
    }
}

func createEssayPageModel(title string, content string) models.WebPage {
    return models.WebPage{
        Title : title,
        MainContent : content,
    }
}
