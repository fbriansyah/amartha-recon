package reconciliation

import (
	"github.com/fbriansyah/amartha-recon/port"
)

type service struct {
	exceptionRepo port.ExceptionRepository
}

func NewService(exceptionRepo port.ExceptionRepository) port.ReconciliationService {
	return &service{
		exceptionRepo: exceptionRepo,
	}
}
