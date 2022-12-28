package http

import (
	"github.com/gofiber/fiber/v2"
	"ojire/app/modules/user/usecase"
	"ojire/config"
)

type UserHttp struct {
	Config      *config.Config
	UserUseCase usecase.InterfaceUseCase
}

func NewHttpUser(e *config.Config, useCase usecase.InterfaceUseCase) InterfaceHttp {
	return &UserHttp{
		Config:      e,
		UserUseCase: useCase,
	}
}

type InterfaceHttp interface {
	MountPublic(app fiber.Router)
	MountPrivate(app fiber.Router)
	RegistrationUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error
	UpdateProfile(c *fiber.Ctx) error
}
