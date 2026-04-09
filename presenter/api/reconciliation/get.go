package reconciliation

import "github.com/gofiber/fiber/v2"

func (h *handler) GetExceptions(c *fiber.Ctx) error {
	exceptions, err := h.reconciliationSvc.GetExceptions()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": exceptions})
}

func (h *handler) GetSuggestions(c *fiber.Ctx) error {
	id := c.Params("id")
	suggestions, err := h.reconciliationSvc.GetSuggestions(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": suggestions})
}
