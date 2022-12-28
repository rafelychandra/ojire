package repositories

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	entityProduct "ojire/app/modules/product/entity"
	"ojire/message"
)

func (psql *PSQLRepository) InsertRelationUserProduct(ctx context.Context, payload entityProduct.RUserProduct) error {
	logCtx := fmt.Sprintf("%T.InsertRelationUserProduct", *psql)

	err := psql.dbClient.Table("r_user_product").Create(&payload).Error
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return err
	}

	return nil
}
