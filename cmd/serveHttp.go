package cmd

import (
	"fmt"
	"log"

	"github.com/fbriansyah/amartha-recon/internal/repository/exception"
	"github.com/fbriansyah/amartha-recon/internal/service/reconciliation"
	healthHandler "github.com/fbriansyah/amartha-recon/presenter/api/health"
	reconapi "github.com/fbriansyah/amartha-recon/presenter/api/reconciliation"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

var serveHttpCmd = &cobra.Command{
	Use:   "serveHttp",
	Short: "Start the HTTP Server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting HTTP Server...")
		
		// Initialize repo and service
		exceptionRepo := exception.NewExceptionRepository() // We instantiate it globally for the server
		reconSvc := reconciliation.NewService(exceptionRepo)

		app := fiber.New()

		// Setup routing
		api := app.Group("/api/v1")
		h := healthHandler.New()
		api.Get("/health", h.HealthCheck)
		
		reconapi.NewHandler(api.Group("/reconciliation"), reconSvc)

		log.Fatal(app.Listen(":8080"))
	},
}

func init() {
	rootCmd.AddCommand(serveHttpCmd)
}
