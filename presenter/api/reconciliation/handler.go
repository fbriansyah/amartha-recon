package reconciliation

import (
	"github.com/fbriansyah/amartha-recon/port"
	"github.com/fbriansyah/amartha-recon/presenter/api"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	api.BaseHandler
	reconciliationSvc port.ReconciliationService
}

func NewHandler(app fiber.Router, reconciliationSvc port.ReconciliationService) {
	h := &handler{
		reconciliationSvc: reconciliationSvc,
	}

	app.Get("/exceptions", h.GetExceptions)
	app.Get("/exceptions/:id/suggestions", h.GetSuggestions)
	app.Post("/exceptions/resolve", h.ResolveException)
	app.Post("/upload/system", h.UploadSystemCSV)
	app.Post("/upload/bank", h.UploadBankCSV)
	app.Post("/process", h.ProcessReconFiles)
}
