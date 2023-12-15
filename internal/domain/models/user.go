package models

type User struct {
	ID          int64
	Email       string
	Gender      string
	Password    []byte
	PhoneNumber string
}
