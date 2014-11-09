package controllers

import (
    "net/http"
    "models"
    "io/ioutil"
)

func RenderEssayTemplate(w http.ResponseWriter, essayTitle string) {
    markdownContent := getEssayMarkdownByEssayTitle(essayTitle)
    essayViewModel := models.WebPage{
        Title : essayTitle,
        MainContent : markdownContent,
    }
    RenderViewModelToTemplate(w, "essayPage", essayViewModel)
}

func getEssayMarkdownByEssayTitle(essayTitle string) string {
    filePath := "content/essays/" + essayTitle + ".md"
    markdownData, err := ioutil.ReadFile(filePath)
    if err != nil {
        panic(err)
    } else {
        return string(markdownData)
    }
}