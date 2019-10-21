package ping

import (
	"fmt"

	"github.com/riari/moni/monzo"
	"github.com/spf13/cobra"
)

// Command supplies the ping command.
func Command(client monzo.Client) *cobra.Command {
	var command = &cobra.Command{
		Use:   "ping",
		Short: "Ping the Monzo API for status info",
	}

	command.AddCommand(
		&cobra.Command{
			Use:   "whoami",
			Short: "Get details of the current authentication",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Print(client.Get("/ping/whoami"))
			},
		},
	)

	return command
}
