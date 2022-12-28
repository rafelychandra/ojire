package repositories

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"ojire/app/modules/product/entity"
	"ojire/message"
)

func (psql *PSQLRepository) InsertProduct(ctx context.Context, payload entity.Product) (uint64, error) {
	logCtx := fmt.Sprintf("%T.InsertProduct", *psql)

	err := psql.dbClient.Table("product").Create(&payload).Error
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return 0, err
	}

	return payload.Id, nil
}

func (psql *PSQLRepository) GetAllProductByUserId(ctx context.Context, userId uint64) ([]entity.Product, error) {
	logCtx := fmt.Sprintf("%T.GetAllProductByUserId", *psql)

	var data []entity.Product
	err := psql.dbClient.Raw(`SELECT p.*
	FROM "user" u
         JOIN r_user_product rup ON rup."userId" = u."id"
         JOIN product p ON p.id = rup."productId"
	WHERE u.id = ?`, userId).Scan(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return nil, err
	}

	return data, nil
}
