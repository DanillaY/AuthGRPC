package models

type App struct {
	ID     int64 `gorm:"primaryKey;autoIncrement:true"`
	Name   string
	Secret string
}
