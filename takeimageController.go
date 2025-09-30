package takeimageController

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"pawscan/entities"
	"pawscan/models/breedModel"
	"pawscan/models/predictionModel"
	"pawscan/models/userModel"
	"pawscan/session"
	"text/template"
)

type PredictionResponse struct {
	Class      uint    `json:"class"` // change from `class_id` to match Python key
	Confidence float32 `json:"confidence"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	sess, err := session.Store.Get(r, session.SessionName)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	userID, ok := sess.Values["userID"]
	if !ok || userID == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	user := entities.MsUser{
		Id:   sess.Values["userID"].(uint),
		Name: sess.Values["userName"].(string),
	}

	tmpl, err := template.ParseFiles("views/Takeimage/index.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, user)
}

func HandlePrediction(w http.ResponseWriter, r *http.Request) {
	// Get session user
	sess, _ := session.Store.Get(r, session.SessionName)
	userID := sess.Values["userID"]

	// Parse image
	file, header, err := r.FormFile("getImg")
	if err != nil {
		http.Error(w, "Failed to read image", http.StatusBadRequest)
		return
	}
	defer file.Close()

	var buf bytes.Buffer
	io.Copy(&buf, file)
	imageBytes := buf.Bytes()

	// Send image to Python
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	part, _ := writer.CreateFormFile("getImg", header.Filename)
	part.Write(imageBytes)
	writer.Close()

	resp, err := http.Post("http://localhost:5000/predict", writer.FormDataContentType(), &requestBody)
	if err != nil {
		http.Error(w, "Prediction server error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var result PredictionResponse
	json.NewDecoder(resp.Body).Decode(&result)

	// Save to DB
	prediction := entities.MsPrediction{
		User:       userModel.GetUserById(userID.(uint)),
		Breed:      breedModel.GetBreedById(result.Class),
		Img:        imageBytes,
		Confidence: result.Confidence,
	}
	predictionModel.InsertPrediction(prediction)

	http.Redirect(w, r, "/result", http.StatusSeeOther)
}
