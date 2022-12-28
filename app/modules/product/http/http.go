package http

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"net/http"
	"ojire/app/modules/product/entity"
	"ojire/message"
	"ojire/utils"
)

func (oh *ProductHttp) Mount(app fiber.Router) {
	app.Post("/insert-product", oh.CreateProduct)
	app.Get("/get-all-product-by-userId", oh.GetAllProduct)
}

func (oh *ProductHttp) CreateProduct(c *fiber.Ctx) error {
	logCtx := fmt.Sprintf("%T.CreateProduct", *oh)

	ctx := context.WithValue(c.Context(), utils.CorrelationIDKey, utils.UUID())

	claims := utils.GetLocalToken(c)
	userId := claims["userId"].(float64)

	var payload entity.Product
	err := c.BodyParser(&payload)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return c.Status(http.StatusBadRequest).JSON(message.RenderResponse(nil, http.StatusBadRequest, false, err.Error()))
	}

	err = oh.ProductUseCase.CreateProduct(ctx, uint64(userId), payload)
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return c.Status(http.StatusInternalServerError).JSON(message.RenderResponse(nil, http.StatusInternalServerError, false, err.Error()))
	}

	message.Log(ctx, logrus.InfoLevel, "success", logCtx)
	return c.Status(http.StatusCreated).JSON(message.RenderResponse(nil, http.StatusCreated, true, "success"))
}

func (oh *ProductHttp) GetAllProduct(c *fiber.Ctx) error {
	logCtx := fmt.Sprintf("%T.GetAllProduct", *oh)

	ctx := context.WithValue(c.Context(), utils.CorrelationIDKey, utils.UUID())

	claims := utils.GetLocalToken(c)
	userId := claims["userId"].(float64)

	output, err := oh.ProductUseCase.GetAllProduct(ctx, uint64(userId))
	if err != nil {
		message.Log(ctx, logrus.ErrorLevel, err.Error(), logCtx)
		return c.Status(http.StatusInternalServerError).JSON(message.RenderResponse(nil, http.StatusInternalServerError, false, err.Error()))
	}

	message.Log(ctx, logrus.InfoLevel, "success", logCtx)
	return c.Status(http.StatusOK).JSON(message.RenderResponse(output, http.StatusOK, true, "success"))
}
