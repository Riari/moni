package account

import (
	"github.com/spf13/cobra"
)

// Command supplies the account command.
func Command() *cobra.Command {
	var command = &cobra.Command{
		Use:   "account",
		Short: "List and select accounts",
	}

	command.AddCommand(
		&cobra.Command{
			Use:   "list",
			Short: "List accounts",
			Run: func(cmd *cobra.Command, args []string) {

			},
		},
		&cobra.Command{
			Use:   "select [account number]",
			Short: "Set a specified account as the active account",
			Long: `Set a specified account (by list number) as the active account.
				Commands that operate on an account will use this account if none is specified.`,
			Args: cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {

			},
		},
	)

	return command
}
