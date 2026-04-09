package reconciliation

import (
	"encoding/csv"
	"os"
	"strings"
	"time"

	"github.com/fbriansyah/amartha-recon/internal/model"
	"github.com/shopspring/decimal"
)

func ParseSystemCSV(filePath string) ([]model.SystemTrx, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// assuming ',' delimiter
	reader.Comma = ','
	
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var results []model.SystemTrx
	// Skip header (trxID,Amount,Type,Date)
	for i, row := range records {
		if i == 0 {
			continue
		}
		if len(row) < 4 {
			continue
		}

		amount, _ := decimal.NewFromString(row[1])
		
		var tType model.TransactionType
		if strings.ToUpper(row[2]) == "CREDIT" {
			tType = model.Credit
		} else {
			tType = model.Debit
		}

		parsedTime, err := time.Parse("2006-01-02 15:04:05", row[3])
		if err != nil {
			parsedTime, _ = time.Parse("2006-01-02", row[3]) // fallback
		}

		results = append(results, model.SystemTrx{
			TrxID:           row[0],
			Amount:          amount,
			Type:            tType,
			TransactionTime: parsedTime,
		})
	}
	return results, nil
}

func ParseBankCSV(filePath string, bankID string) ([]model.BankStatement, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// assuming ',' delimiter
	reader.Comma = ','
	
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var results []model.BankStatement
	// Skip header (ReferenceNo,Amount,Date)
	for i, row := range records {
		if i == 0 {
			continue
		}
		if len(row) < 3 {
			continue
		}

		amount, _ := decimal.NewFromString(row[1])
		parsedTime, _ := time.Parse("2006-01-02", strings.TrimSpace(row[2]))

		results = append(results, model.BankStatement{
			BankID:           bankID,
			UniqueIdentifier: row[0],
			Amount:           amount,
			Date:             parsedTime,
		})
	}
	return results, nil
}
