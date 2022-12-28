package utils

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const CorrelationIDKey = "CorrelationID"

func GetCorrelationIDFromContext(ctx context.Context) string {
	id, ok := ctx.Value(CorrelationIDKey).(string)
	if !ok {
		return ""
	}
	return id
}

func UUID() string {
	return uuid.New().String()
}

func GetLocalToken(c *fiber.Ctx) jwt.MapClaims {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims
}
