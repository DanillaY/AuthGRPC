package models

type User struct {
	ID          int64 `gorm:"primaryKey;autoIncrement:true"`
	Email       string
	Gender      string
	Password    []byte
	PhoneNumber string
}
