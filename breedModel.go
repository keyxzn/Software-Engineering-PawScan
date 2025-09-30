package breedModel

import (
	"pawscan/config"
	"pawscan/entities"
)

func GetBreedById(breedID uint) entities.MsBreed {
	query := `
				SELECT
					BreedId,
					BreedName,
					BreedDescription,
					OriginId,
					SizeId,
					TypeId
				FROM MsBreed
				WHERE BreedId = $1
	`
	var breed entities.MsBreed

	err := config.DB.QueryRow(query, breedID).Scan(&breed.Id, &breed.Name, &breed.Description, &breed.Origin.Id, &breed.Size.Id, &breed.Type.Id)
	if err != nil {
		panic(err)
	}

	return breed
}
