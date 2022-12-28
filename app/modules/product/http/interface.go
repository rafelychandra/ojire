package http

import (
	"github.com/gofiber/fiber/v2"
	"ojire/app/modules/product/usecase"
	"ojire/config"
)

type ProductHttp struct {
	Config         *config.Config
	ProductUseCase usecase.InterfaceUseCase
}

func NewHttpProduct(e *config.Config, productUseCase usecase.InterfaceUseCase) InterfaceHttp {
	return &ProductHttp{
		Config:         e,
		ProductUseCase: productUseCase,
	}
}

type InterfaceHttp interface {
	Mount(app fiber.Router)
	CreateProduct(c *fiber.Ctx) error
	GetAllProduct(c *fiber.Ctx) error
}
