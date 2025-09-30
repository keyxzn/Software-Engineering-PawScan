package hasilController

import (
	"encoding/base64"
	"net/http"
	"pawscan/models/predictionModel"
	"pawscan/session"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// Get user session
	sess, err := session.Store.Get(r, session.SessionName)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Type assert userID
	userID, ok := sess.Values["userID"].(uint)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Call your existing model function
	result := predictionModel.GetLatestPredictionByUserID(userID)

	// Prepare data for template
	data := map[string]interface{}{
		"BreedName":   result.Breed.Name,
		"Origin":      result.Breed.Origin.Name,
		"Type":        result.Breed.Type.Name,
		"Size":        result.Breed.Size.Name,
		"Description": result.Breed.Description,
		"Confidence":  result.Confidence * 100,
		"ImageBase64": "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(result.Img),
	}

	// Parse and execute template
	tmpl, err := template.ParseFiles("views/hasil/index.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}