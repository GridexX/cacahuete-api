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
