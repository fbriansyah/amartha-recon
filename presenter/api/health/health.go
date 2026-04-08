package healthHandler

import (
	"github.com/fbriansyah/amartha-recon/internal/model/common"
	"github.com/gofiber/fiber/v2"
)

func (h *HealthCheckHandler) HealthCheck(c *fiber.Ctx) error {
	return c.JSON(common.NewResponse().Success(nil))
}
