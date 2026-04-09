package reconciliation

import (
	"path/filepath"
	"strings"
	"time"

	reconmodel "github.com/fbriansyah/amartha-recon/internal/model/recon"
	"github.com/fbriansyah/amartha-recon/internal/service/reconciliation"
	"github.com/gofiber/fiber/v2"
)

type ProcessRequest struct {
	StartDate string `json:"start_date"` // YYYY-MM-DD
	EndDate   string `json:"end_date"`   // YYYY-MM-DD
}

func (h *handler) ProcessReconFiles(c *fiber.Ctx) error {
	var req ProcessRequest
	if err := c.BodyParser(&req); err != nil {
		return h.SendError(c, fiber.StatusBadRequest, "invalid payload")
	}

	startData, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return h.SendError(c, fiber.StatusBadRequest, "invalid start_date format, expected YYYY-MM-DD")
	}
	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return h.SendError(c, fiber.StatusBadRequest, "invalid end_date format, expected YYYY-MM-DD")
	}
	// make enddate cover the whole day
	endDate = endDate.Add(24 * time.Hour).Add(-time.Second)

	sysDir := "storage/recon-files/system"
	bnkDir := "storage/recon-files/banks"

	sysFiles, _ := filepath.Glob(filepath.Join(sysDir, "*.csv"))
	bnkFiles, _ := filepath.Glob(filepath.Join(bnkDir, "*.csv"))

	var systemTrxs []reconmodel.SystemTrx
	var bankStatements []reconmodel.BankStatement

	// Extract System CSVs bounded by date string or inside the CSV itself
	for _, sf := range sysFiles {
		trxs, err := reconciliation.ParseSystemCSV(sf)
		if err == nil {
			for _, t := range trxs {
				if t.TransactionTime.After(startData.Add(-1*time.Second)) && t.TransactionTime.Before(endDate) {
					systemTrxs = append(systemTrxs, t)
				}
			}
		}
	}

	for _, bf := range bnkFiles {
		// extract bank id visually
		filename := filepath.Base(bf)
		bankID := strings.Split(filename, "_")[0]
		statements, err := reconciliation.ParseBankCSV(bf, bankID)
		if err == nil {
			for _, s := range statements {
				if s.Date.After(startData.Add(-1*time.Second)) && s.Date.Before(endDate) {
					bankStatements = append(bankStatements, s)
				}
			}
		}
	}

	result, err := h.reconciliationSvc.Reconcile(systemTrxs, bankStatements)
	if err != nil {
		return h.SendError(c, fiber.StatusInternalServerError, err.Error())
	}

	// Group missing bank exceptions
	bankGroups := make(map[string][]reconmodel.ExceptionRecord)
	for _, b := range result.BankExceptions {
		// we encoded raw string so we can parse it back, or just use BankID if it was available on root
		// but since only ID is there we fall back to generic mapping, wait, raw json has BankID
		bankGroups["BANK"] = append(bankGroups["BANK"], b)
	}

	return h.SendSuccess(c, fiber.Map{
		"total_processed":   result.TotalProcessed,
		"total_matched":     result.TotalMatched,
		"total_unmatched":   result.TotalUnmatched,
		"total_discrepancy": result.TotalDiscrepancy.String(),
		"details": fiber.Map{
			"missing_in_bank":   result.SystemExceptions,
			"missing_in_system": bankGroups, // currently just grouped into BANK key unless json parsed
		},
	})
}
