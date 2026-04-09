package exception

import "errors"

func (r *ExceptionRepository) UpdateStatus(id string, status string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	e, ok := r.exceptions[id]
	if !ok {
		return errors.New("exception not found")
	}

	e.Status = status
	r.exceptions[id] = e
	return nil
}
