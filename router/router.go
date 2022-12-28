package router

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/sirupsen/logrus"
	"net/http"
	"ojire/message"
	"ojire/utils"
)

func (hr *HandlerRouter) ListRouter() *fiber.App {
	app := fiber.New()

	hr.Setup.HttpUser.MountPublic(app)
	meAPI := app.Group("/me", hr.AuthRequired())
	hr.Setup.HttpUser.MountPrivate(meAPI)

	api := app.Group("/product", hr.AuthRequired())
	hr.Setup.HttpProduct.Mount(api)

	return app
}

func (hr *HandlerRouter) AuthRequired() func(ctx *fiber.Ctx) error {
	logCtx := fmt.Sprintf("%T.AuthRequired", *hr)
	return jwtware.New(jwtware.Config{
		SuccessHandler: func(c *fiber.Ctx) error {
			ctx := context.WithValue(c.Context(), utils.CorrelationIDKey, utils.UUID())

			claims := utils.GetLocalToken(c)
			if claims == nil {
				err := errors.New("invalid user")
				message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
				return c.Status(http.StatusUnauthorized).JSON(message.RenderResponse(nil, http.StatusUnauthorized, false, err.Error()))
			}

			email := claims["email"].(string)
			password := claims["password"].(string)

			data, err := hr.Setup.UnitOfWork.PSQLRepository.GetUserByEmail(ctx, email)
			if err != nil {
				message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
				return c.Status(http.StatusUnauthorized).JSON(message.RenderResponse(nil, http.StatusUnauthorized, false, err.Error()))
			}

			if data.Password != password {
				err = errors.New("invalid user")
				message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
				return c.Status(http.StatusUnauthorized).JSON(message.RenderResponse(nil, http.StatusUnauthorized, false, err.Error()))
			}

			err = c.Next()
			if err != nil {
				message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
				return err
			}
			return nil
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			message.Log(nil, logrus.ErrorLevel, err.Error(), logCtx)
			return c.Status(http.StatusUnauthorized).JSON(message.RenderResponse(nil, http.StatusUnauthorized, false, err.Error()))
		},
		SigningKey: []byte(hr.Setup.Env.SecretKey),
	})
}
