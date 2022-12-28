package usecase

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	entityProduct "ojire/app/modules/product/entity"
	"ojire/message"
)

func (pu *ProductUseCase) CreateProduct(ctx context.Context, userId uint64, payload entityProduct.Product) error {
	logCtx := fmt.Sprintf("%T.CreateProduct", *pu)

	sku, err := entityProduct.GenerateSKU(payload.Quantity)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return err
	}
	payload.SKU = sku
	tx := pu.UnitOfWork.Start()

	productId, err := tx.PSQLRepository.InsertProduct(ctx, payload)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		_ = tx.Dispose()
		return err
	}

	err = tx.PSQLRepository.InsertRelationUserProduct(ctx, entityProduct.RUserProduct{
		UserId:    userId,
		ProductId: productId,
	})
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		_ = tx.Dispose()
		return err
	}

	err = tx.Complete()
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return err
	}

	return nil
}

func (pu *ProductUseCase) GetAllProduct(ctx context.Context, userId uint64) ([]entityProduct.Product, error) {
	logCtx := fmt.Sprintf("%T.GetAllProduct", *pu)

	data, err := pu.UnitOfWork.PSQLRepository.GetAllProductByUserId(ctx, userId)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return nil, err
	}

	return data, nil
}
