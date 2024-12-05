package team

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func NewSet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set <team>",
		Short: "Sets your team",
		Long:  `Set the team which is used to identify your timer among others`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) != 1 {
				log.Fatal("Command requires 'team' to be included as argument")
			}

		},
	}

	return cmd
}
