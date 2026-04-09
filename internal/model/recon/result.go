package reconmodel

import "github.com/shopspring/decimal"

type ReconciliationResult struct {
	TotalProcessed   int
	TotalMatched     int
	TotalUnmatched   int
	SystemExceptions []ExceptionRecord
	BankExceptions   []ExceptionRecord
	TotalDiscrepancy decimal.Decimal
}
