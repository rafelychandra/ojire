package usecase

import (
	"context"
	"ojire/app/modules/user/entity"
	"ojire/app/repositories"
	"ojire/config"
)

type UserUseCase struct {
	Config     *config.Config
	UnitOfWork *repositories.UnitOfWork
}

func NewUseCaseUser(e *config.Config, unitOfWork *repositories.UnitOfWork) InterfaceUseCase {
	return &UserUseCase{
		Config:     e,
		UnitOfWork: unitOfWork,
	}
}

type InterfaceUseCase interface {
	RegistrationUser(ctx context.Context, payload entity.User) error
	Login(ctx context.Context, payload entity.User) (string, error)
	ChangePassword(ctx context.Context, email, password string) error
	UpdateProfile(ctx context.Context, email string, payload entity.User) error
}
