package controllers

import (
        "net/http"
        "appengine"
        "appengine/mail"
        "fmt"
)

func HandleEmailInquiry(w http.ResponseWriter, r *http.Request) {
    appengineContext := appengine.NewContext(r)
    senderEmail := r.FormValue("email")
    emailContent := r.FormValue("content")
    msg := &mail.Message{
        Sender:  "qdonnellan@gmail.com",
        ReplyTo: senderEmail,
        To:      []string{"quentin@qdonnellan.com"},
        Subject: "INQUIRY: Email from qdonnellan.com/contact",
        Body:    emailContent,
    }
    if err := mail.Send(appengineContext, msg); err != nil {
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprint(w, JsonResponse{"success": false, "message": err})
        appengineContext.Errorf("Couldn't send email: %v", err)
        return
    } else {
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprint(w, JsonResponse{"success": true})
        return
    }
}

