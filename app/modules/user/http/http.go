package http

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"net/http"
	"ojire/app/modules/user/entity"
	"ojire/message"
	"ojire/utils"
)

func (oh *UserHttp) MountPublic(app fiber.Router) {
	app.Post("/registration-user", oh.RegistrationUser)
	app.Post("/login", oh.Login)
}

func (oh *UserHttp) MountPrivate(app fiber.Router) {
	app.Put("/change-password", oh.ChangePassword)
	app.Put("/update-profile", oh.UpdateProfile)
}

func (oh *UserHttp) RegistrationUser(c *fiber.Ctx) error {
	logCtx := fmt.Sprintf("%T.RegistrationUser", *oh)

	ctx := context.WithValue(c.Context(), utils.CorrelationIDKey, utils.UUID())

	var payload entity.User
	err := c.BodyParser(&payload)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return c.Status(http.StatusBadRequest).JSON(message.RenderResponse(nil, http.StatusBadRequest, false, err.Error()))
	}

	err = oh.UserUseCase.RegistrationUser(ctx, payload)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return c.Status(http.StatusInternalServerError).JSON(message.RenderResponse(nil, http.StatusInternalServerError, false, err.Error()))
	}

	message.Log(ctx, logrus.InfoLevel, "success", logCtx)
	return c.Status(http.StatusCreated).JSON(message.RenderResponse(nil, http.StatusCreated, true, "success"))
}

func (oh *UserHttp) Login(c *fiber.Ctx) error {
	logCtx := fmt.Sprintf("%T.Login", *oh)

	ctx := context.WithValue(c.Context(), utils.CorrelationIDKey, utils.UUID())

	var payload entity.User
	err := c.BodyParser(&payload)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return c.Status(http.StatusBadRequest).JSON(message.RenderResponse(nil, http.StatusBadRequest, false, err.Error()))
	}

	token, err := oh.UserUseCase.Login(ctx, payload)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return c.Status(http.StatusInternalServerError).JSON(message.RenderResponse(nil, http.StatusInternalServerError, false, err.Error()))
	}

	message.Log(ctx, logrus.InfoLevel, "success", logCtx)
	return c.Status(http.StatusOK).JSON(message.RenderResponse(token, http.StatusOK, true, "success"))
}

func (oh *UserHttp) ChangePassword(c *fiber.Ctx) error {
	logCtx := fmt.Sprintf("%T.ChangePassword", *oh)

	ctx := context.WithValue(c.Context(), utils.CorrelationIDKey, utils.UUID())

	claims := utils.GetLocalToken(c)
	email := claims["email"].(string)

	var payload entity.User
	err := c.BodyParser(&payload)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return c.Status(http.StatusBadRequest).JSON(message.RenderResponse(nil, http.StatusBadRequest, false, err.Error()))
	}

	err = oh.UserUseCase.ChangePassword(ctx, email, payload.Password)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return c.Status(http.StatusInternalServerError).JSON(message.RenderResponse(nil, http.StatusInternalServerError, false, err.Error()))
	}

	message.Log(ctx, logrus.InfoLevel, "success", logCtx)
	return c.Status(http.StatusOK).JSON(message.RenderResponse(nil, http.StatusOK, true, "success"))
}

func (oh *UserHttp) UpdateProfile(c *fiber.Ctx) error {
	logCtx := fmt.Sprintf("%T.UpdateProfile", *oh)

	ctx := context.WithValue(c.Context(), utils.CorrelationIDKey, utils.UUID())

	claims := utils.GetLocalToken(c)
	email := claims["email"].(string)

	var payload entity.User
	err := c.BodyParser(&payload)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return c.Status(http.StatusBadRequest).JSON(message.RenderResponse(nil, http.StatusBadRequest, false, err.Error()))
	}

	err = oh.UserUseCase.UpdateProfile(ctx, email, payload)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return c.Status(http.StatusInternalServerError).JSON(message.RenderResponse(nil, http.StatusInternalServerError, false, err.Error()))
	}

	message.Log(ctx, logrus.InfoLevel, "success", logCtx)
	return c.Status(http.StatusOK).JSON(message.RenderResponse(nil, http.StatusOK, true, "success"))
}
