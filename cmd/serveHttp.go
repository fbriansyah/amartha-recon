package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serveHttpCmd = &cobra.Command{
	Use:   "serveHttp",
	Short: "Serve HTTP",
	Long:  `Serve HTTP`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serveHttp")
	},
}

func init() {
	rootCmd.AddCommand(serveHttpCmd)
}
