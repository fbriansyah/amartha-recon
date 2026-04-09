package port

import reconmodel "github.com/fbriansyah/amartha-recon/internal/model/recon"

type ReconciliationService interface {
	Reconcile(systemTrx []reconmodel.SystemTrx, bankStatements []reconmodel.BankStatement) (*reconmodel.ReconciliationResult, error)
	GetExceptions() ([]reconmodel.ExceptionRecord, error)
	GetSuggestions(exceptionID string) ([]reconmodel.SystemTrx, error)
	ResolveException(exceptionID string, action string, systemTrxID *string) error
}
