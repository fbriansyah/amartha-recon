package exception

import (
	"errors"

	reconmodel "github.com/fbriansyah/amartha-recon/internal/model/recon"
)

func (r *ExceptionRepository) FindExceptions() ([]reconmodel.ExceptionRecord, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []reconmodel.ExceptionRecord
	for _, e := range r.exceptions {
		result = append(result, e)
	}
	return result, nil
}

func (r *ExceptionRepository) FindExceptionByID(id string) (*reconmodel.ExceptionRecord, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if e, ok := r.exceptions[id]; ok {
		return &e, nil
	}
	return nil, errors.New("exception not found")
}

func (r *ExceptionRepository) FindSystemTrxByAmount(amount string) ([]reconmodel.SystemTrx, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []reconmodel.SystemTrx
	for _, trx := range r.systemTrx {
		if trx.Amount.String() == amount {
			result = append(result, trx)
		}
	}
	return result, nil
}
