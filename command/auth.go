package command

import (
	"fmt"

	"github.com/riari/moni/monzo"
	"github.com/spf13/cobra"
)

// Auth supplies the auth command.
func Auth(client monzo.Client) *cobra.Command {
	var command = &cobra.Command{
		Use:   "auth",
		Short: "Authentication commands",
	}

	command.AddCommand(
		&cobra.Command{
			Use:   "status",
			Short: "Get details of the current authentication",
			Run: func(cmd *cobra.Command, args []string) {
				data := client.Auth.GetStatus()

				if data.Error != "" {
					fmt.Printf("%s: %s", data.Code, data.ErrorDescription)
					return
				}

				if data.Authenticated {
					fmt.Printf("Authenticated as %s", data.UserID)
				} else {
					fmt.Print("Unauthenticated")
				}
			},
		},
	)

	return command
}
