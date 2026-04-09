package reconmodel

import (
	"time"

	"github.com/shopspring/decimal"
)

type TransactionType string

const (
	Debit  TransactionType = "DEBIT"
	Credit TransactionType = "CREDIT"
)

// Internal System Data
type SystemTrx struct {
	TrxID           string
	Amount          decimal.Decimal
	Type            TransactionType
	TransactionTime time.Time
}

// External Bank Data
type BankStatement struct {
	BankID           string
	UniqueIdentifier string
	Amount           decimal.Decimal
	Date             time.Time
}

// Normalized Data for the Matcher
type NormalizedRecord struct {
	SourceID  string
	BankID    string
	AmountAbs decimal.Decimal
	Type      TransactionType
	Date      time.Time
}
