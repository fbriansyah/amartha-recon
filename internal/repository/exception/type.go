package exception

import (
	"sync"

	"github.com/fbriansyah/amartha-recon/internal/model"
)

type ExceptionRepository struct {
	exceptions map[string]model.ExceptionRecord
	systemTrx  []model.SystemTrx
	mu         sync.RWMutex
}

func NewExceptionRepository() *ExceptionRepository {
	return &ExceptionRepository{
		exceptions: make(map[string]model.ExceptionRecord),
		systemTrx:  make([]model.SystemTrx, 0),
	}
}
