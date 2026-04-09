package port

import "github.com/fbriansyah/amartha-recon/internal/model"

type ExceptionRepository interface {
	SaveExceptions(exceptions []model.ExceptionRecord) error
	FindExceptions() ([]model.ExceptionRecord, error)
	FindExceptionByID(id string) (*model.ExceptionRecord, error)
	UpdateStatus(id string, status string) error
	
	// Helper functions to fetch potential suggestions
	FindSystemTrxByAmount(amount string) ([]model.SystemTrx, error)
	SaveSystemTrx(trx []model.SystemTrx) error
}
