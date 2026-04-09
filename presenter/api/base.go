package api

import (
	"github.com/fbriansyah/amartha-recon/internal/model/common"
	"github.com/gofiber/fiber/v2"
)

type BaseHandler struct{}

func (b *BaseHandler) SendError(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(common.NewResponse().Error(code, message))
}

func (b *BaseHandler) SendSuccess(c *fiber.Ctx, data any) error {
	return c.JSON(common.NewResponse().Success(data))
}
