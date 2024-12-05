package team

import (
	"fmt"

	"github.com/fehlhabers/zt/internal/core/config"
	"github.com/spf13/cobra"
)

func NewList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all configured teams",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			zt, err := config.ZtConfig()
			if err != nil {
				return
			}

			fmt.Printf("Active team: %s\n\n", zt.ActiveTeam)
			fmt.Printf("Teams:\n")
			for name := range zt.Teams {
				fmt.Printf("- %s\n", name)
			}
		},
	}

	return cmd
}

