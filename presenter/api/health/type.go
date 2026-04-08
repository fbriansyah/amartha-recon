package healthHandler

import (
	"github.com/fbriansyah/amartha-recon/port"
)

type HealthCheckHandler struct{}

func New() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

var _ port.IHealthCheckHandler = (*HealthCheckHandler)(nil)
