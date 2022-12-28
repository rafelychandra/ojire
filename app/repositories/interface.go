package repositories

import (
	"context"
	"github.com/jinzhu/gorm"
	entityProduct "ojire/app/modules/product/entity"
	entityUser "ojire/app/modules/user/entity"
	"ojire/config"
)

type UnitOfWork struct {
	Config         *config.Config
	PSQLRepository InterfacePSQLRepository
	Tx             *gorm.DB
}

func NewUnitOfWork(e *config.Config, dbClient *gorm.DB) *UnitOfWork {
	return &UnitOfWork{
		Config:         e,
		PSQLRepository: NewPSQLRepository(dbClient),
		Tx:             dbClient,
	}
}

type InterfaceTransaction interface {
	Start() *UnitOfWork
	Complete() error
	Dispose() error
}

type PSQLRepository struct {
	dbClient *gorm.DB
}

func NewPSQLRepository(dbClient *gorm.DB) InterfacePSQLRepository {
	return &PSQLRepository{
		dbClient: dbClient,
	}
}

type InterfacePSQLRepository interface {
	Ping(ctx context.Context) error
	InsertUser(ctx context.Context, payload entityUser.User) error
	GetUserByEmail(ctx context.Context, email string) (*entityUser.User, error)
	UpdatePassword(ctx context.Context, payload entityUser.User) error
	UpdateDataUser(ctx context.Context, payload entityUser.User) error
	InsertProduct(ctx context.Context, payload entityProduct.Product) (uint64, error)
	InsertRelationUserProduct(ctx context.Context, payload entityProduct.RUserProduct) error
	GetAllProductByUserId(ctx context.Context, userId uint64) ([]entityProduct.Product, error)
}
