package port

import "github.com/gofiber/fiber/v2"

type IHealthCheckHandler interface {
	HealthCheck(c *fiber.Ctx) error
}
