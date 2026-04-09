package port

import "github.com/fbriansyah/amartha-recon/internal/model"

type ReconciliationService interface {
	Reconcile(systemTrx []model.SystemTrx, bankStatements []model.BankStatement) (*model.ReconciliationResult, error)
	GetExceptions() ([]model.ExceptionRecord, error)
	GetSuggestions(exceptionID string) ([]model.SystemTrx, error)
	ResolveException(exceptionID string, action string, systemTrxID *string) error
}
