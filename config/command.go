package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Command supplies the config command.
func Command() *cobra.Command {
	var command = &cobra.Command{
		Use:   "config",
		Short: "Get or set config values",
	}

	command.AddCommand(
		&cobra.Command{
			Use:   "get [key]",
			Short: "Get a config value by key",
			Args:  cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Print(viper.GetString(args[0]))
			},
		},
		&cobra.Command{
			Use:   "set [key] [value]",
			Short: "Set a config value by key",
			Args:  cobra.MinimumNArgs(2),
			Run: func(cmd *cobra.Command, args []string) {
				viper.Set(args[0], args[1])
				viper.WriteConfig()
				fmt.Printf("Config key %s set to %s", args[0], args[1])
			},
		},
	)

	return command
}
