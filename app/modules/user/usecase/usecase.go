package usecase

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"ojire/app/modules/user/entity"
	"ojire/message"
	"time"

	"fmt"
)

func (ou *UserUseCase) RegistrationUser(ctx context.Context, payload entity.User) error {
	logCtx := fmt.Sprintf("%T.RegistrationUser", *ou)

	data, err := ou.UnitOfWork.PSQLRepository.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return err
	}

	if data != nil && payload.Email == data.Email {
		return errors.New("email sudah terdaftar")
	}

	err = payload.Validation()
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return err
	}

	hashPassword, err := payload.HashPassword()
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return err
	}

	payload.Password = hashPassword
	err = ou.UnitOfWork.PSQLRepository.InsertUser(ctx, payload)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return err
	}

	return nil
}

func (ou *UserUseCase) Login(ctx context.Context, payload entity.User) (string, error) {
	logCtx := fmt.Sprintf("%T.Login", *ou)

	data, err := ou.UnitOfWork.PSQLRepository.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return "", err
	}

	if data == nil {
		err = errors.New("user not found")
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return "", err
	}

	if !payload.CheckPasswordHash(data.Password) {
		err = errors.New("password anda salah")
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return "", err
	}

	token, err := ou.generateToken(ctx, data.Id, data.Name, data.Email, data.Password)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return "", err
	}

	return token, nil
}

func (ou *UserUseCase) ChangePassword(ctx context.Context, email, password string) error {
	logCtx := fmt.Sprintf("%T.ChangePassword", *ou)

	payload := entity.User{
		Email:    email,
		Password: password,
	}

	hashPassword, err := payload.HashPassword()
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return err
	}

	payload.Password = hashPassword
	payload.UpdatedAt = time.Now()
	err = ou.UnitOfWork.PSQLRepository.UpdatePassword(ctx, payload)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return err
	}

	return nil
}

func (ou *UserUseCase) UpdateProfile(ctx context.Context, email string, payload entity.User) error {
	logCtx := fmt.Sprintf("%T.UpdateProfile", *ou)

	payload.Email = email
	err := payload.Validation()
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return err
	}

	err = ou.UnitOfWork.PSQLRepository.UpdateDataUser(ctx, payload)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return err
	}

	return nil
}
