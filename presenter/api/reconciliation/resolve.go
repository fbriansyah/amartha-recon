package reconciliation

import "github.com/gofiber/fiber/v2"

func (h *handler) ResolveException(c *fiber.Ctx) error {
	type ResolveRequest struct {
		ExceptionID string  `json:"exception_id"`
		Action      string  `json:"action"` // "FORCE_MATCH", "RETURN", "WRITE_OFF"
		SystemTrxID *string `json:"system_trx_id"`
	}

	var req ResolveRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}

	if err := h.reconciliationSvc.ResolveException(req.ExceptionID, req.Action, req.SystemTrxID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "success"})
}
