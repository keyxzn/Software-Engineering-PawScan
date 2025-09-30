package authController

import (
	"net/http"
	"pawscan/session"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess, err := session.Store.Get(r, session.SessionName)
	if err != nil {
		http.Error(w, "Unable to get session", http.StatusInternalServerError)
		return
	}

	// Invalidate the session
	sess.Options.MaxAge = -1
	err = sess.Save(r, w)
	if err != nil {
		http.Error(w, "Unable to clear session", http.StatusInternalServerError)
		return
	}

	// Redirect to login page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}