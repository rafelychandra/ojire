package usecase

import (
	"context"
	entityProduct "ojire/app/modules/product/entity"
	"ojire/app/repositories"
	"ojire/config"
)

type ProductUseCase struct {
	Config     *config.Config
	UnitOfWork *repositories.UnitOfWork
}

func NewUseCaseProduct(e *config.Config, unitOfWork *repositories.UnitOfWork) InterfaceUseCase {
	return &ProductUseCase{
		Config:     e,
		UnitOfWork: unitOfWork,
	}
}

type InterfaceUseCase interface {
	CreateProduct(ctx context.Context, userId uint64, payload entityProduct.Product) error
	GetAllProduct(ctx context.Context, userId uint64) ([]entityProduct.Product, error)
}
