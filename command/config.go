package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config supplies the config command.
func Config(conf *viper.Viper) *cobra.Command {
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
				fmt.Print(conf.GetString(args[0]))
			},
		},
		&cobra.Command{
			Use:   "set [key] [value]",
			Short: "Set a config value by key",
			Args:  cobra.MinimumNArgs(2),
			Run: func(cmd *cobra.Command, args []string) {
				conf.Set(args[0], args[1])
				conf.WriteConfig()
				fmt.Printf("Config key %s set", args[0])
			},
		},
	)

	return command
}
