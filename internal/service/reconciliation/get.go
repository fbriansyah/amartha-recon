package reconciliation

import reconmodel "github.com/fbriansyah/amartha-recon/internal/model/recon"

func (s *service) GetExceptions() ([]reconmodel.ExceptionRecord, error) {
	return s.exceptionRepo.FindExceptions()
}

func (s *service) GetSuggestions(exceptionID string) ([]reconmodel.SystemTrx, error) {
	e, err := s.exceptionRepo.FindExceptionByID(exceptionID)
	if err != nil {
		return nil, err
	}
	return s.exceptionRepo.FindSystemTrxByAmount(e.Amount.String())
}
