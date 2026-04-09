package exception

import "github.com/fbriansyah/amartha-recon/internal/model"

func (r *ExceptionRepository) SaveExceptions(exceptions []model.ExceptionRecord) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, e := range exceptions {
		r.exceptions[e.ID] = e
	}
	return nil
}

func (r *ExceptionRepository) SaveSystemTrx(trx []model.SystemTrx) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.systemTrx = append(r.systemTrx, trx...)
	return nil
}
