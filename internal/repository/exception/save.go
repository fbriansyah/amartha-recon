package exception

import reconmodel "github.com/fbriansyah/amartha-recon/internal/model/recon"

func (r *ExceptionRepository) SaveExceptions(exceptions []reconmodel.ExceptionRecord) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, e := range exceptions {
		r.exceptions[e.ID] = e
	}
	return nil
}

func (r *ExceptionRepository) SaveSystemTrx(trx []reconmodel.SystemTrx) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.systemTrx = append(r.systemTrx, trx...)
	return nil
}
