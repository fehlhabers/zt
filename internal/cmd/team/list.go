package team

import (
	"fmt"

	"github.com/fehlhabers/zt/internal/global"
	"github.com/spf13/cobra"
)

func NewList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all configured teams",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {

			state := global.GetStateKeeper().GetState()

			fmt.Printf("Active Team: %s\n\n", state.TeamName)
			fmt.Printf("Teams:\n")
			for name := range state.AllTeams {
				fmt.Printf("- %s\n", name)
			}
		},
	}

	return cmd
}
