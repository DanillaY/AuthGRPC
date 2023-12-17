package storage

import (
	"context"
	"errors"
	"log/slog"
	"main/internal/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PGDatabase struct {
	logger slog.Logger
	db     *gorm.DB
}

func New(logger slog.Logger, host string, user string, password string, port string) (db *PGDatabase, err error) {
	var base PGDatabase
	//dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=postgres" + " port=" + port
	//TO DO fixport
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=Users" + " port=5432" + " sslmode=disable"
	pg, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Could not connect to database")
	}
	base.logger = logger
	base.db = pg
	base.db.AutoMigrate(&models.App{}, &models.User{})
	return &base, nil
}

func (pg *PGDatabase) App(ctx context.Context, appID int64) (app models.App, err error) {
	var appD models.App
	pg.db.Where("ID = ?", appID).First(&appD)
	if appD.ID == 0 {
		return models.App{}, errors.New("No such app")
	}
	return appD, nil
}

func (pg *PGDatabase) CreateUser(ctx context.Context, email string, gender string, phoneNumber string, passHash []byte) (id int64, err error) {
	user := models.User{Email: email, Password: passHash, Gender: gender, PhoneNumber: phoneNumber}
	result := pg.db.Create(&user)

	if result.Error != nil {
		pg.logger.Error("Could not add new user in the database")
		return -1, result.Error
	}
	return user.ID, nil
}

func (pg *PGDatabase) ValidateUser(ctx context.Context, email string) (user models.User, err error) {
	var userE models.User
	pg.db.Where("Email = ?", email).First(&userE)
	if userE.ID == 0 {
		return models.User{}, errors.New("Could not find user with such credentials")
	}
	return userE, nil
}
