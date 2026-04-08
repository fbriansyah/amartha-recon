package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "amartha-recon",
	Short: "Amartha Recon is a tool for reconciling financial data",
	Long:  `Amartha Recon is a tool for reconciling financial data`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
