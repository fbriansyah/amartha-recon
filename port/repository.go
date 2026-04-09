package port

import reconmodel "github.com/fbriansyah/amartha-recon/internal/model/recon"

type ExceptionRepository interface {
	SaveExceptions(exceptions []reconmodel.ExceptionRecord) error
	FindExceptions() ([]reconmodel.ExceptionRecord, error)
	FindExceptionByID(id string) (*reconmodel.ExceptionRecord, error)
	UpdateStatus(id string, status string) error

	// Helper functions to fetch potential suggestions
	FindSystemTrxByAmount(amount string) ([]reconmodel.SystemTrx, error)
	SaveSystemTrx(trx []reconmodel.SystemTrx) error
}
