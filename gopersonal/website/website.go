package website

import (
    "net/http"
    "html/template"
)

type Page struct {
    Title string
    TemplateName string
}

// Handle the front page of the website
func FrontPageHandler(w http.ResponseWriter, r *http.Request) {
    frontPage := Page{
        Title : "FrontPage",
        TemplateName : "front_page.html", 
    }
    renderPage(w, frontPage)
}

func renderPage(w http.ResponseWriter, page Page) {
    pageTemplate := loadAndReturnTemplate(page.TemplateName)
    pageTemplate.Execute(w, page)
}

func loadAndReturnTemplate(templateName string) *template.Template {
    templateDir := "gopersonal/website/templates/"
    templatePath := templateDir + templateName
    newTemplate := template.Must(template.ParseFiles(templatePath))
    return newTemplate
}