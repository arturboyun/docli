package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "docli",
	Long: `Useful cli for Docker.`,
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	return nil
}
