package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var profile string

var rootCmd = &cobra.Command{
	Use: "literalog",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&profile, "profile", "dev", "default profile is dev")
}
