package loginController

import (
	"fmt"
	"html/template"
	"net/http"
	"pawscan/models/userModel"
	"pawscan/session"

	"github.com/gorilla/sessions"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/Login/index.html") // FIXED PATH
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Template parse error:", err)
		return
	}
	tmpl.Execute(w, nil)
}

func CheckLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("== CheckLogin called ==")
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	// 1. Check if user exists
	user, err := userModel.GetUserByEmail(email)
	if err != nil {
		tmpl, err := template.ParseFiles("views/Login/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, map[string]interface{}{
			"Alert": "Email is incorrect. Please recheck. If you don't have an account, please sign in.",
		})
		return
	}

	// 2. Check password
	if user.Password != password {
		tmpl, err := template.ParseFiles("views/Login/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, map[string]interface{}{
			"Alert": "Password is incorrect. Please recheck.",
		})
		return
	}

	// 3. Create session and redirect
	sessionData, _ := session.Store.Get(r, session.SessionName)
	sessionData.Values["userID"] = user.Id
	sessionData.Values["userName"] = user.Name

	sessionData.Options = &sessions.Options{
    Path:     "/",
    MaxAge:   3600,
    HttpOnly: true,
	}
	sessionData.Save(r, w)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
