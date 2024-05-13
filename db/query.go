package db

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var loger = logrus.WithFields(logrus.Fields{
	"context": "db/query",
})

func LogAndReturnError(l *logrus.Entry, result *gorm.DB, action string, modelType string) error {
	if err := result.Error; err != nil {
		l.WithError(err).Error("Error when trying to query database to " + action + " " + modelType)
		return err
	}
	return nil
}

func CreateUser(db *gorm.DB, user User) (*User, error) {

	// TODO Here we need to test if the user was already created
	result := db.Where("username = ? OR email = ?", user.FirstName, user.Email).FirstOrCreate(&user)
	db.Create(&user)
	err := LogAndReturnError(loger, result, "create", "user")
	return &user, err
}

func GetUsername(db *gorm.DB, username string) (*User, error) {
	user := new(User)
	result := db.Where("username = ?", username).First(user)
	err := LogAndReturnError(loger, result, "get", "username")
	return user, err
}

func UpsertToken(db *gorm.DB, token Token) (*Token, error) {

	tokenR := Token{
		UserID: token.UserID,
	}
	result := db.Where("user_id = ?", token.UserID).Assign(Token{Value: token.Value, ExpirationDate: token.ExpirationDate}).FirstOrCreate(&tokenR)
	err := LogAndReturnError(loger, result, "upsert", "token")
	return &tokenR, err
}

func GetTokenUser(db *gorm.DB, userID uint) (*Token, error) {
	tokenR := new(Token)
	result := db.Where("user_id = ?", userID).First(tokenR)
	err := LogAndReturnError(loger, result, "get", "token username")
	return tokenR, err
}

func DeleteToken(db *gorm.DB, userID uint) error {
	result := db.Where("user_id = ?", userID).Delete(&Token{})
	err := LogAndReturnError(loger, result, "delete", "token")
	return err
}
