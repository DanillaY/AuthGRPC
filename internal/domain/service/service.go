package auth

import (
	"context"
	"log/slog"
	"main/internal/domain/jwt"
	"main/internal/domain/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	log           *slog.Logger
	tokenLifeTime time.Duration
	base          IDatabase
	app           IApp
}

type IDatabase interface {
	CreateUser(ctx context.Context, email string, passHash []byte) (id int64, err error)
	ValidateUser(ctx context.Context, email string) (user models.User, err error)
}
type IApp interface {
	App(ctx context.Context, appID int64) (app models.App, err error)
}

func New(logger *slog.Logger, tokenLifetime time.Duration, base IDatabase, app IApp) *Auth {

	return &Auth{
		log:           logger,
		tokenLifeTime: tokenLifetime,
		base:          base,
		app:           app,
	}
}

func (a *Auth) Login(ctx context.Context, email string, gender string, password string, phoneNumber string, appID int64) (token string, err error) {
	a.log.Info("Loggin a user")
	user, err := a.base.ValidateUser(ctx, email)
	if err != nil {
		a.log.Error("No such user ")
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		a.log.Error("Wrong credentials")
	}
	app, err := a.app.App(ctx, appID)
	if err != nil {
		a.log.Error("No such app")
	}

	jwtToken, err := jwt.GenNewToken(user, app, a.tokenLifeTime)
	if err != nil {
		a.log.Error("Could not generate new token")
	}
	return jwtToken, nil

}
func (a *Auth) Register(ctx context.Context, email string, gender string, phoneNumber string, password string) (userID int64, err error) {
	a.log.Info("Registering a user")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		a.log.Error("Could not create hash")
	}

	id, err := a.base.CreateUser(ctx, email, hash)
	if err != nil {
		a.log.Error("Could not register new user")
	}
	return id, nil
}
