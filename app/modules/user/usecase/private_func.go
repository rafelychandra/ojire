package usecase

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"ojire/message"
	"time"
)

func (ou *UserUseCase) generateToken(ctx context.Context, id uint64, name, email, password string) (string, error) {
	logCtx := fmt.Sprintf("%T.generateToken", *ou)

	claims := jwt.MapClaims{
		"userId":   id,
		"name":     name,
		"email":    email,
		"password": password,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(ou.Config.SecretKey))
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return "", err
	}

	return tokenString, nil
}
