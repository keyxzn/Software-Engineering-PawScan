package entities

type MsPrediction struct {
	Id uint
	User MsUser
	Breed MsBreed
	Img []byte
	Confidence float32
}