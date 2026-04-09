package reconciliation

import "github.com/gofiber/fiber/v2"

func (h *handler) GetExceptions(c *fiber.Ctx) error {
	exceptions, err := h.reconciliationSvc.GetExceptions()
	if err != nil {
		return h.SendError(c, fiber.StatusInternalServerError, err.Error())
	}
	return h.SendSuccess(c, exceptions)
}

func (h *handler) GetSuggestions(c *fiber.Ctx) error {
	id := c.Params("id")
	suggestions, err := h.reconciliationSvc.GetSuggestions(id)
	if err != nil {
		return h.SendError(c, fiber.StatusInternalServerError, err.Error())
	}
	return h.SendSuccess(c, suggestions)
}
