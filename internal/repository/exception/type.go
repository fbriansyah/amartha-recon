package exception

import (
	"sync"

	reconmodel "github.com/fbriansyah/amartha-recon/internal/model/recon"
)

type ExceptionRepository struct {
	exceptions map[string]reconmodel.ExceptionRecord
	systemTrx  []reconmodel.SystemTrx
	mu         sync.RWMutex
}

func NewExceptionRepository() *ExceptionRepository {
	return &ExceptionRepository{
		exceptions: make(map[string]reconmodel.ExceptionRecord),
		systemTrx:  make([]reconmodel.SystemTrx, 0),
	}
}
