package homeController

import (
	"net/http"
	"pawscan/entities"
	"pawscan/session"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess, err := session.Store.Get(r, session.SessionName)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Check if user is logged in
	userID, ok := sess.Values["userID"]
	if !ok || userID == nil {
		// Not logged in: redirect to login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Logged in: load home page with user info
	tmpl, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	data := entities.MsUser{
		Id:    sess.Values["userID"].(uint),
		Name:  sess.Values["userName"].(string),
	}
	tmpl.Execute(w, data)
}
