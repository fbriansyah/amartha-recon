package reconciliation

import (
	"encoding/json"
	"sort"
	"time"

	"github.com/fbriansyah/amartha-recon/internal/model"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// getLookbackDates helper function generates acceptable lookup dates
// T+1, and accounts for weekend gaps (e.g. Monday bank statement can look up to Friday)
func getLookbackDates(bankDate time.Time) []time.Time {
	dates := []time.Time{bankDate}
	if bankDate.Weekday() == time.Monday {
		dates = append(dates, bankDate.AddDate(0, 0, -1)) // Sun
		dates = append(dates, bankDate.AddDate(0, 0, -2)) // Sat
		dates = append(dates, bankDate.AddDate(0, 0, -3)) // Fri
	} else {
		dates = append(dates, bankDate.AddDate(0, 0, -1)) // Previous day
	}

	// Sort dates from oldest to newest (FIFO requirement)
	sort.SliceStable(dates, func(i, j int) bool {
		return dates[i].Before(dates[j])
	})

	return dates
}

func (s *service) Reconcile(systemTrx []model.SystemTrx, bankStatements []model.BankStatement) (*model.ReconciliationResult, error) {
	// Group System records by Date -> Amount (String format for exact matching) -> List (Bucket)
	buckets := make(map[string]map[string][]model.SystemTrx)
	tfmt := "2006-01-02"

	for _, st := range systemTrx {
		dateKey := st.TransactionTime.Format(tfmt)
		amtKey := st.Amount.String()

		if _, ok := buckets[dateKey]; !ok {
			buckets[dateKey] = make(map[string][]model.SystemTrx)
		}
		buckets[dateKey][amtKey] = append(buckets[dateKey][amtKey], st)
	}

	// Sort buckets by specific transaction time to enforce FIFO
	for dateKey := range buckets {
		for amtKey := range buckets[dateKey] {
			bucket := buckets[dateKey][amtKey]
			sort.SliceStable(bucket, func(i, j int) bool {
				return bucket[i].TransactionTime.Before(bucket[j].TransactionTime)
			})
			buckets[dateKey][amtKey] = bucket
		}
	}

	var sysExceptions []model.ExceptionRecord
	var bnkExceptions []model.ExceptionRecord
	bankTolerance := decimal.NewFromInt(5000)

	totalProcessed := len(systemTrx) + len(bankStatements)
	totalMatched := 0
	totalDiscrepancy := decimal.NewFromInt(0)

	// Iterate through Bank Statements
	for _, bs := range bankStatements {
		lookbackDates := getLookbackDates(bs.Date)
		matched := false

		for _, lookupDate := range lookbackDates {
			if matched {
				break
			}
			dateKey := lookupDate.Format(tfmt)
			dayBuckets, ok := buckets[dateKey]
			if !ok {
				continue
			}

			amtKey := bs.Amount.String()
			bucketQueue := dayBuckets[amtKey]
			if len(bucketQueue) > 0 {
				// Matched with oldest exact amount
				buckets[dateKey][amtKey] = bucketQueue[1:]
				matched = true
				totalMatched += 2 // 1 bank, 1 system match
				break
			} else {
				// Advanced: Check for tolerance variations across all queues for this date
				for stAmtKey, stQueue := range dayBuckets {
					if len(stQueue) == 0 {
						continue
					}
					stAmt, _ := decimal.NewFromString(stAmtKey)
					diff := bs.Amount.Sub(stAmt).Abs()
					if diff.LessThanOrEqual(bankTolerance) {
						// Matched with oldest under tolerance
						buckets[dateKey][stAmtKey] = stQueue[1:]
						matched = true
						totalMatched += 2
						totalDiscrepancy = totalDiscrepancy.Add(diff)
						break
					}
				}
			}
		}

		if !matched {
			raw, _ := json.Marshal(bs)
			bnkExceptions = append(bnkExceptions, model.ExceptionRecord{
				ID:           uuid.NewString(), // assuming we use uuid
				Source:       "BANK",
				OriginalDate: bs.Date,
				Amount:       bs.Amount,
				Type:         model.Credit, // just default for PoC structure
				RawData:      string(raw),
				Status:       "OPEN",
			})
		}
	}

	// Residual System Transactions -> Exceptions
	for _, dayBuckets := range buckets {
		for _, queue := range dayBuckets {
			for _, st := range queue {
				raw, _ := json.Marshal(st)
				sysExceptions = append(sysExceptions, model.ExceptionRecord{
					ID:           uuid.NewString(),
					Source:       "SYSTEM",
					OriginalDate: st.TransactionTime,
					Amount:       st.Amount,
					Type:         st.Type,
					RawData:      string(raw),
					Status:       "OPEN",
				})
			}
		}
	}

	allExceptions := append(sysExceptions, bnkExceptions...)
	s.exceptionRepo.SaveSystemTrx(systemTrx)
	s.exceptionRepo.SaveExceptions(allExceptions)

	res := &model.ReconciliationResult{
		TotalProcessed:   totalProcessed,
		TotalMatched:     totalMatched,
		TotalUnmatched:   len(allExceptions),
		SystemExceptions: sysExceptions,
		BankExceptions:   bnkExceptions,
		TotalDiscrepancy: totalDiscrepancy,
	}

	return res, nil
}
