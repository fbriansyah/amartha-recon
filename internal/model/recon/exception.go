package reconmodel

import (
	"time"

	"github.com/shopspring/decimal"
)

type ExceptionRecord struct {
	ID           string          `json:"id"`
	Source       string          `json:"source"` // "SYSTEM" or "BANK"
	OriginalDate time.Time       `json:"original_date"`
	Amount       decimal.Decimal `json:"amount"`
	Type         TransactionType `json:"type"`
	RawData      string          `json:"raw_data"` // JSON dump of CSV row (or system object)
	Status       string          `json:"status"`   // "OPEN", "RESOLVED"
}
