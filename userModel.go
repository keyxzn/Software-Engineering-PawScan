package userModel

import (
	"pawscan/config"
	"pawscan/entities"
)

func GetUserByEmail(email string) (*entities.MsUser, error) {
	var user entities.MsUser
	err := config.DB.QueryRow("SELECT UserID, UserName, UserPassword FROM MsUser WHERE UserEmail = $1", email).
		Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserById(userID uint) entities.MsUser {
	var user entities.MsUser
	err := config.DB.QueryRow("Select UserID, UserName, UserEmail, UserPassword FROM MsUser WHERE UserID = $1", userID).
		Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		panic(err)
	}

	return user
}

func IsEmailExist(email string) bool {
	query := `SELECT COUNT(*) FROM msUser WHERE UserEmail = $1`
	var count int
	err := config.DB.QueryRow(query, email).Scan(&count)
	return err == nil && count > 0
}

func InsertUser(user entities.MsUser) uint {
	query := `
		INSERT INTO msUser(UserName, UserEmail, UserPassword)
		VALUES ($1, $2, $3) RETURNING UserID
	`

	var newID uint
	err := config.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&newID)
	if err != nil {
		return 0
	}

	return newID
}