package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/beatrizrdgs/literalog/internal/initializer"
	"github.com/beatrizrdgs/literalog/internal/server"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		if profile != "prod" {
			err := godotenv.Load(".env")
			if err != nil {
				return err
			}
		}

		err := initializer.InitServices()
		if err != nil {
			return err
		}

		handlers := initializer.InitHandlers()

		s := server.NewServer("1234", handlers)
		s.Start()

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		log.Println("Shutting down server...")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
