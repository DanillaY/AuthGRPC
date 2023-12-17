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
	Log           *slog.Logger
	TokenLifeTime time.Duration
	Base          IDatabase
	App           IApp
}

type IDatabase interface {
	CreateUser(ctx context.Context, email string, gender string, phoneNumber string, passHash []byte) (id int64, err error)
	ValidateUser(ctx context.Context, email string) (user models.User, err error)
}
type IApp interface {
	App(ctx context.Context, appID int64) (app models.App, err error)
}

func New(logger *slog.Logger, tokenLifetime time.Duration, base IDatabase, app IApp) *Auth {

	return &Auth{
		Log:           logger,
		TokenLifeTime: tokenLifetime,
		Base:          base,
		App:           app,
	}
}

func (a *Auth) Login(ctx context.Context, email string, gender string, password string, phoneNumber string, appID int64) (token string, err error) {
	a.Log.Info("Loggin a user")
	user, err := a.Base.ValidateUser(ctx, email)
	if err != nil {
		a.Log.Error("No such user ")
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		a.Log.Error("Wrong credentials")
	}
	app, err := a.App.App(ctx, appID)
	if err != nil {
		a.Log.Error("No such app")
	}

	jwtToken, err := jwt.GenNewToken(user, app, a.TokenLifeTime)
	if err != nil {
		a.Log.Error("Could not generate new token")
	}
	return jwtToken, nil

}
func (a *Auth) Register(ctx context.Context, email string, gender string, phoneNumber string, password string) (userID int64, err error) {
	a.Log.Info("Registering a user")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		a.Log.Error("Could not create hash")
	}

	id, err := a.Base.CreateUser(ctx, email, gender, phoneNumber, hash)
	if err != nil {
		a.Log.Error("Could not register new user")
	}
	return id, nil
}
