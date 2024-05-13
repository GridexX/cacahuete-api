package db

type Travel struct {
	ID uint `gorm:"primaryKey;autoIncrement:true;uniqueIndex;not null"`
}

type User struct {
	ID       uint `gorm:"primaryKey;autoIncrement:true;uniqueIndex;not null"`
	email    string
	username string
	
}
