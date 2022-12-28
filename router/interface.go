package router

import (
	"ojire/service"

	"github.com/gofiber/fiber/v2"
)

type HandlerRouter struct {
	Setup *service.HandlerSetup
}

func NewHandlerRouter(setup *service.HandlerSetup) InterfaceRouter {
	return &HandlerRouter{
		Setup: setup,
	}
}

type InterfaceRouter interface {
	ListRouter() *fiber.App
	AuthRequired() func(ctx *fiber.Ctx) error
}
