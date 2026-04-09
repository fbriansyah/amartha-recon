package reconciliation

import "github.com/fbriansyah/amartha-recon/internal/model"

func (s *service) GetExceptions() ([]model.ExceptionRecord, error) {
	return s.exceptionRepo.FindExceptions()
}

func (s *service) GetSuggestions(exceptionID string) ([]model.SystemTrx, error) {
	e, err := s.exceptionRepo.FindExceptionByID(exceptionID)
	if err != nil {
		return nil, err
	}
	return s.exceptionRepo.FindSystemTrxByAmount(e.Amount.String())
}
