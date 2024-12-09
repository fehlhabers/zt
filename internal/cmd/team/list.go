package team

import (
	"fmt"

	"github.com/fehlhabers/zt/internal/adapter/state/config"
	"github.com/spf13/cobra"
)

func NewList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all configured teams",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			teams := config.ListTeams()

			fmt.Printf("Teams:\n")
			for name := range teams {
				fmt.Printf("- %s\n", name)
			}
		},
	}

	return cmd
}
