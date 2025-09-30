package historyController

import (
	"encoding/base64"
	"html/template"
	"net/http"
	"pawscan/entities"
	"pawscan/models/predictionModel"
	"pawscan/session"
)

type PredictionView struct {
	entities.MsPrediction
	Base64Image string
}

func Index(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess, err := session.Store.Get(r, session.SessionName)
	if err != nil {
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	userID, ok := sess.Values["userID"]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Get all predictions
	rawPredictions := predictionModel.GetAllPredictionFromUserId(userID.(uint))

	// Wrap predictions with Base64Image
	var predictions []PredictionView
	for _, p := range rawPredictions {
		predictions = append(predictions, PredictionView{
			MsPrediction: p,
			Base64Image:  base64.StdEncoding.EncodeToString(p.Img),
		})
	}

	// Parse and render template
	tmpl, err := template.ParseFiles("views/History/index.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, predictions)
}