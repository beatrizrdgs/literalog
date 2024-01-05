package cmd

import (
	"github.com/spf13/cobra"
)

var profile string

var rootCmd = &cobra.Command{
	Use:   "library",
	Short: "literalog's management of books and related entities",
}

func init() {
	rootCmd.PersistentFlags().StringVar(&profile, "profile", "development", "application profile (default is development)")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
