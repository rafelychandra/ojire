package repositories

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"ojire/message"
)

func (psql *PSQLRepository) Ping(ctx context.Context) error {
	logCtx := fmt.Sprintf("%T.Ping", *psql)

	err := psql.dbClient.DB().Ping()
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return err
	}

	return nil
}
