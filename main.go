package main

import (
	"log"
	"net/http"
	"pawscan/config"
	"pawscan/controllers/authController"
	"pawscan/controllers/hasilController"
	"pawscan/controllers/historyController"
	"pawscan/controllers/homeController"
	"pawscan/controllers/landingPageController"
	"pawscan/controllers/loginController"
	"pawscan/controllers/signinController"
	"pawscan/controllers/takeimageController"
	"pawscan/session"

	"github.com/gorilla/sessions"
)

func main() {
	session.Store = sessions.NewCookieStore([]byte("super-secret-key"))
	session.SessionName = "pawscan-session"

	config.ConnectDB()
	log.Println("Connection is established, continue to opening port")

	// 1. Landing Page
	http.HandleFunc("/", landingPageController.Welcome)

	// 2. Login Page
	http.HandleFunc("/login", loginController.Index)
	http.HandleFunc("/checklogin", loginController.CheckLogin)
	
	// 3. Signin Page
	http.HandleFunc("/signin", signinController.Index)
	http.HandleFunc("/checksignin", signinController.SignIn)

	// 4. home Page
	http.HandleFunc("/home", homeController.Index)

	// 5. Take Image
	http.HandleFunc("/takeimage", takeimageController.Index)
	http.HandleFunc("/predict", takeimageController.HandlePrediction)

	// 5. Hasil
	http.HandleFunc("/result", hasilController.Index)

	// 6. History
	http.HandleFunc("/history", historyController.Index)

	// 7. Logout
	http.HandleFunc("/logout", authController.Logout)



	// handle css and js
	http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("views"))))

	log.Println("Server running on http://localhost:8080")
	log.Println("Other url can be found on http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}