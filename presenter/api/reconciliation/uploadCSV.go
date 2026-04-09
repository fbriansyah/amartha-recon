package reconciliation

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) UploadSystemCSV(c *fiber.Ctx) error {
	return h.handleUpload(c, "system")
}

func (h *handler) UploadBankCSV(c *fiber.Ctx) error {
	return h.handleUpload(c, "bank")
}

func (h *handler) handleUpload(c *fiber.Ctx, source string) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse file upload: " + err.Error()})
	}

	// Ensure the directory exists
	dir := "storage/recon-files"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create directory: " + err.Error()})
	}

	// Create unique file name
	filename := fmt.Sprintf("%s/%s-%d-%s", dir, source, time.Now().Unix(), file.Filename)

	// Save file to destination
	if err := c.SaveFile(file, filename); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file: " + err.Error()})
	}

	return c.JSON(fiber.Map{
		"message":  fmt.Sprintf("%s file uploaded successfully", source),
		"filename": filename,
	})
}
