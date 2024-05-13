package db

import "time"

type Journey struct {
	ID            uint `gorm:"primaryKey;autoIncrement:true;uniqueIndex;not null"`
	TrainId       uint
	DepartureDate time.Time
	ArrivalDate   time.Time
}

type User struct {
	ID         uint `gorm:"primaryKey;autoIncrement:true;uniqueIndex;not null"`
	Email      string
	Username   string
	Password   string
	FirstName  string
	LastName   string
	Street     string
	PostalCode string
	City       string
	Journeys   []Journey `gorm:"many2many:user_journeys;"`
}

type Station struct {
	ID          uint `gorm:"primaryKey;autoIncrement:true;uniqueIndex;not null"`
	StationName string
	City        string
	PostalCode  uint
	Insee       uint
}

type Token struct {
	ID             uint `gorm:"primaryKey;autoIncrement:true;uniqueIndex;not null"`
	Value          string
	ExpirationDate time.Time
}
