package jwt

import (
	"main/internal/domain/models"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenNewToken(user models.User, app models.App, duration time.Duration) (token string, err error) {
	genToken := jwt.New(jwt.SigningMethodHS256)
	claims := genToken.Claims.(jwt.MapClaims)
	claims["userID"] = user.ID
	claims["email"] = user.Email
	claims["expDuration"] = time.Now().Add(duration).Unix()
	claims["appID"] = app.ID

	stringToken, err := genToken.SignedString([]byte(app.Secret))

	if err != nil {
		return "", err
	}
	return stringToken, nil
}
