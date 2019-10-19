package main

import (
	"github.com/riari/moni/account"
	"github.com/riari/moni/config"
	"github.com/spf13/cobra"
)

func main() {
	config.Initialise()

	var rootCmd = &cobra.Command{Use: "moni"}
	rootCmd.AddCommand(account.Command())
	rootCmd.AddCommand(config.Command())

	rootCmd.Execute()
}
