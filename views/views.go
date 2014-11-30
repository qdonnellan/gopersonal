package views

import (
    "net/http"
    "github/go-personal/controllers"
)

// FrontPageView handles request for the / page.
func FrontPageView(w http.ResponseWriter, r *http.Request) {
    controllers.RenderBasicPage(w, "frontPage") 
}

// ContactPageView handles request for the /contact page. 
func ContactPageView(w http.ResponseWriter, r *http.Request) {
    controllers.RenderBasicPage(w, "contactPage")
}

// AboutPageView handles request for the /about page. 
func AboutPageView(w http.ResponseWriter, r *http.Request) {
    controllers.RenderBasicPage(w, "aboutPage")
}

// ConstructionPageView handles requests to show the construction page
func ConstructionPageView(w http.ResponseWriter, r *http.Request) {
    controllers.RenderBasicPage(w, "constructionPage")
}

// ProjectPageView handles request for the /project page. 
func ProjectPageView(w http.ResponseWriter, r *http.Request) {
    controllers.RenderBasicPage(w, "projectPage")
}

// EssayPageView handles requests matching /essays/ prefix or shows the
// essay list page if just /essays/
func EssayPageView(w http.ResponseWriter, r *http.Request) {
    // skip the first characters: /essays/ (length 8)
    essayTitle := r.URL.Path[8:]
    if len(essayTitle) == 0 {
        controllers.RenderBasicPage(w, "essayListPage")
    } else {
        controllers.RenderEssayPage(w, essayTitle)
    }
}

// StaticFileViewHandler handles all /static/* requests. 
func StaticFileViewHandler(w http.ResponseWriter, r *http.Request) {
    staticFileName := r.URL.Path[1:]
    http.ServeFile(w, r, staticFileName)
}



