package reconciliation

func (s *service) ResolveException(exceptionID string, action string, systemTrxID *string) error {
	status := "RESOLVED"
	if action == "RETURN" {
		status = "RETURNED"
	}
	return s.exceptionRepo.UpdateStatus(exceptionID, status)
}
