package signinController

import (
	"net/http"
	"pawscan/entities"
	"pawscan/models/userModel"
	"pawscan/session"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/Signin/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}

	// Get form data
	email := r.FormValue("email")
	fullName := r.FormValue("full-name")
	password := r.FormValue("password")

	// Check if email already exists
	if userModel.IsEmailExist(email) {
		http.Redirect(w, r, "/signin?error=email_exists", http.StatusSeeOther)
		return
	}

	// Create user object
	user := entities.MsUser{
		Name:     fullName,
		Email:    email,
		Password: password,
	}

	// Insert into DB
	newID := userModel.InsertUser(user)
	if newID == 0 {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Set session
	sess, _ := session.Store.Get(r, session.SessionName)
	sess.Values["userID"] = user.Id
	sess.Values["userName"] = user.Name
	sess.Save(r, w)

	// Redirect to homepage
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}