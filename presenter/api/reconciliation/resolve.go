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
		return h.SendError(c, fiber.StatusBadRequest, "invalid payload")
	}

	if err := h.reconciliationSvc.ResolveException(req.ExceptionID, req.Action, req.SystemTrxID); err != nil {
		return h.SendError(c, fiber.StatusInternalServerError, err.Error())
	}

	return h.SendSuccess(c, nil)
}
