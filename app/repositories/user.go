package repositories

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"ojire/app/modules/user/entity"
	"ojire/message"
)

func (psql *PSQLRepository) InsertUser(ctx context.Context, payload entity.User) error {
	logCtx := fmt.Sprintf("%T.InsertUser", *psql)

	err := psql.dbClient.Table("user").Create(&payload).Error
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return err
	}

	return nil
}

func (psql *PSQLRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	logCtx := fmt.Sprintf("%T.GetUserByEmail", *psql)

	var data entity.User
	err := psql.dbClient.Table("user").Where(`email = ?`, email).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return nil, err
	}

	return &data, nil
}

func (psql *PSQLRepository) UpdatePassword(ctx context.Context, payload entity.User) error {
	logCtx := fmt.Sprintf("%T.UpdatePassword", *psql)

	err := psql.dbClient.Table("user").Where(`email = ?`, payload.Email).Updates(&payload).Error
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return err
	}

	return nil
}

func (psql *PSQLRepository) UpdateDataUser(ctx context.Context, payload entity.User) error {
	logCtx := fmt.Sprintf("%T.UpdateDataUser", *psql)

	err := psql.dbClient.Table("user").Where(`email = ?`, payload.Email).Updates(payload).Error
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return err
	}
	return nil
}
