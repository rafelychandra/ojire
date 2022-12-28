package service

import (
	"context"
	"ojire/app/repositories"
	"os"

	httpProductPkg "ojire/app/modules/product/http"
	useCaseProductPkg "ojire/app/modules/product/usecase"
	httpUserPkg "ojire/app/modules/user/http"
	useCaseUserPkg "ojire/app/modules/user/usecase"
	"ojire/config"
	"ojire/infrastructure"
	"ojire/message"

	"github.com/sirupsen/logrus"
)

type HandlerSetup struct {
	Env             *config.Config
	EnvironmentName *string
	Http            *bool
	DatabaseClient  *infrastructure.HandlerDatabase
	UnitOfWork      *repositories.UnitOfWork
	HttpUser        httpUserPkg.InterfaceHttp
	HttpProduct     httpProductPkg.InterfaceHttp
}

func MakeHandler(ctx context.Context, Environment *string, Http *bool) *HandlerSetup {
	loadConfig := config.HandlerLoadConfig{
		Env: *Environment,
	}

	env, err := loadConfig.LoadConfig()
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), "SETUP CONFIG")
		os.Exit(1)
	}

	databaseClient, err := infrastructure.NewDatabaseClient(env)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), "SETUP DATABASE CLIENT")
		os.Exit(1)
	}

	unitOfWork := repositories.NewUnitOfWork(env, databaseClient.DatabaseClient)
	useCaseUser := useCaseUserPkg.NewUseCaseUser(env, unitOfWork)
	httpUser := httpUserPkg.NewHttpUser(env, useCaseUser)
	useCaseProduct := useCaseProductPkg.NewUseCaseProduct(env, unitOfWork)
	httpProduct := httpProductPkg.NewHttpProduct(env, useCaseProduct)

	return &HandlerSetup{
		Env:             env,
		EnvironmentName: Environment,
		Http:            Http,
		DatabaseClient:  databaseClient,
		UnitOfWork:      unitOfWork,
		HttpUser:        httpUser,
		HttpProduct:     httpProduct,
	}
}
