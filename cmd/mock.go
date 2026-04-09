package cmd

import (
	"fmt"
	"time"

	"github.com/fbriansyah/amartha-recon/internal/model"
	"github.com/fbriansyah/amartha-recon/internal/repository/exception"
	"github.com/fbriansyah/amartha-recon/internal/service/reconciliation"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

var mockCmd = &cobra.Command{
	Use:   "mock",
	Short: "Run reconciliation service with mock data",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Mock Data Generation...")

		repo := exception.NewExceptionRepository()
		svc := reconciliation.NewService(repo)

		// Set up explicit dates.
		// Monday, April 13, 2026
		monday := time.Date(2026, time.April, 13, 12, 0, 0, 0, time.UTC)
		// previous Friday, April 10, 2026
		friday := time.Date(2026, time.April, 10, 15, 0, 0, 0, time.UTC)
		// previous Sunday, April 12, 2026
		sunday := time.Date(2026, time.April, 12, 12, 0, 0, 0, time.UTC)
		
		// Create dummy data
		systemTrx := []model.SystemTrx{
			{
				TrxID:           "SYS-001",
				Amount:          decimal.NewFromInt(100000),
				Type:            model.Credit,
				TransactionTime: friday, // Friday
			},
			{
				TrxID:           "SYS-002",
				Amount:          decimal.NewFromInt(50000),
				Type:            model.Debit,
				TransactionTime: sunday, // Sunday
			},
			{
				TrxID:           "SYS-003",
				Amount:          decimal.NewFromInt(72000), // Won't match perfectly, has difference > 5000 from Bank
				Type:            model.Debit,
				TransactionTime: sunday,
			},
		}

		bankStatements := []model.BankStatement{
			{
				BankID:           "BNK-001",
				UniqueIdentifier: "REF-001",
				Amount:           decimal.NewFromInt(100000),
				Date:             monday, // Monday matching against Friday SYS-001
			},
			{
				BankID:           "BNK-002",
				UniqueIdentifier: "REF-002",
				Amount:           decimal.NewFromInt(48000), // within 5000 tolerance of SYS-002 (50000)
				Date:             monday,
			},
			{
				BankID:           "BNK-003",
				UniqueIdentifier: "REF-003",
				Amount:           decimal.NewFromInt(90000), // No matching SYS record
				Date:             monday,
			},
		}


		result, err := svc.Reconcile(systemTrx, bankStatements)
		if err != nil {
			fmt.Printf("Error reconciling: %v\n", err)
			return
		}

		fmt.Printf("Reconciliation complete. Processed: %d, Matched: %d, Unmatched: %d\n", result.TotalProcessed, result.TotalMatched, result.TotalUnmatched)
		fmt.Printf("Total Discrepancy Amount: %s\n", result.TotalDiscrepancy.String())
		for _, ex := range append(result.SystemExceptions, result.BankExceptions...) {
			fmt.Printf("Exception [%s]: Amount %s -> Source: %s\n", ex.ID, ex.Amount.String(), ex.Source)
		}
	},
}

func init() {
	rootCmd.AddCommand(mockCmd)
}
