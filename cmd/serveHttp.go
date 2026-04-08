package cmd

import (
	"fmt"

	healthHandler "github.com/fbriansyah/amartha-recon/presenter/api/health"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

var serveHttpCmd = &cobra.Command{
	Use:   "serveHttp",
	Short: "Serve HTTP",
	Long:  `Serve HTTP`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting HTTP Server...")
		app := fiber.New()

		healthHandler := healthHandler.New()

		// routes
		app.Get("/health", healthHandler.HealthCheck)

		if err := app.Listen(":8080"); err != nil {
			fmt.Printf("Error starting server: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveHttpCmd)
}
