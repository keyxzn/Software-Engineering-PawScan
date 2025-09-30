package predictionModel

import (
	"pawscan/config"
	"pawscan/entities"
)

func InsertPrediction(prediction entities.MsPrediction) bool{
	query := `
				INSERT INTO MsPrediction(UserID, BreedID, PredictionImage, PredictionConfidence)
				VALUES ($1, $2, $3, $4) RETURNING PredictionId
		`
	var result uint
	err := config.DB.QueryRow(query, prediction.User.Id, prediction.Breed.Id, prediction.Img, prediction.Confidence).Scan(&result)
	if err != nil {
		panic(err)
	}

	return result > 0
}


func GetLatestPredictionByUserID(userID uint) entities.MsPrediction {
	query := `
				SELECT 
					mp.predictionId,
					mp.predictionImage,
					mp.predictionConfidence,
					mu.UserId, 
					mu.UserName, 
					mu.UserEmail, 
					mu.UserPassword,
					mb.BreedId, 
					mb.BreedName, 
					mo.OriginId, 
					mo.OriginName,
					ms.SizeId, 
					ms.SizeName,
					mt.TypeId, 
					mt.TypeName,
					mb.BreedDescription
				FROM 
					MsPrediction mp
				JOIN 
					msUser mu ON mu.UserID = mp.UserID
				JOIN 
					MsBreed mb ON mb.BreedID = mp.BreedId
				JOIN
					MsSize ms ON ms.SizeID = mb.SizeID
				JOIN
					MsOrigin mo ON mo.OriginId = mb.OriginID
				JOIN
					MsType mt ON mt.TypeID = mb.TypeID
				WHERE 
					mu.UserID = $1
				ORDER BY 
					mp.PredictionID DESC
				LIMIT 1
			`

	row := config.DB.QueryRow(query, userID)

	var prediction entities.MsPrediction
	var imgBytes []byte

	err := row.Scan(
		&prediction.Id,
		&imgBytes,
		&prediction.Confidence,
		&prediction.User.Id, &prediction.User.Name, &prediction.User.Email, &prediction.User.Password,
		&prediction.Breed.Id, &prediction.Breed.Name,
		&prediction.Breed.Origin.Id, &prediction.Breed.Origin.Name,
		&prediction.Breed.Size.Id, &prediction.Breed.Size.Name,
		&prediction.Breed.Type.Id, &prediction.Breed.Type.Name,
		&prediction.Breed.Description,
	)

	if err != nil {
		panic(err)
	}

	// Assign byte image to struct
	prediction.Img = imgBytes

	return prediction
}

func GetAllPredictionFromUserId(userID uint) []entities.MsPrediction {	
	query := `
				SELECT 
					mp.predictionId,
					mp.predictionImage,
					mp.predictionConfidence,
					mu.UserId, 
					mu.UserName, 
					mu.UserEmail, 
					mu.UserPassword,
					mb.BreedId, 
					mb.BreedName, 
					mo.OriginId, 
					mo.OriginName,
					ms.SizeId, 
					ms.SizeName,
					mt.TypeId, 
					mt.TypeName,
					mb.BreedDescription
					FROM 
					MsPrediction mp
				JOIN 
					msUser mu ON mu.UserID = mp.UserID
				JOIN 
					MsBreed mb ON mb.BreedID = mp.BreedId
				JOIN
					MsSize ms ON ms.SizeID = mb.SizeID
				JOIN
					MsOrigin mo ON mo.OriginId = mb.OriginID
				JOIN
					MsType mt ON mt.TypeID = mb.TypeID
				WHERE 
					mu.UserID = $1
				ORDER BY 
					mp.PredictionID ASC
			`
		
		rows, err := config.DB.Query(query, userID)
		if err != nil {
			panic(err)
		}
		
		defer rows.Close()
		
		var predictions []entities.MsPrediction
		var imgBytes []byte


		for rows.Next() {
			var prediction entities.MsPrediction
			err := rows.Scan(
				&prediction.Id,
				&imgBytes,
				&prediction.Confidence,
				&prediction.User.Id, &prediction.User.Name, &prediction.User.Email, &prediction.User.Password,
				&prediction.Breed.Id, &prediction.Breed.Name,
				&prediction.Breed.Origin.Id, &prediction.Breed.Origin.Name,
				&prediction.Breed.Size.Id, &prediction.Breed.Size.Name,
				&prediction.Breed.Type.Id, &prediction.Breed.Type.Name,
				&prediction.Breed.Description,
			)
			if err != nil {
				panic(err)
			}
			prediction.Img = imgBytes

			predictions = append(predictions, prediction)
		}


	return predictions
}