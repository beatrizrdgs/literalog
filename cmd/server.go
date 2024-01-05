package cmd

import (
	"github.com/literalog/library/internal/app/gateways/api"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "start",
	Short: "starts library",
	RunE: func(cmd *cobra.Command, args []string) error {
		server := api.NewServer(":8080")
		return server.ServeHttp()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
