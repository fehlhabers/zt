package team

import (
	"github.com/fehlhabers/zt/internal/core/team"
	"github.com/spf13/cobra"
)

func NewConfigure() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Configure a team",
		Long:  `Configure a new team with setting using a wizard`,
		Run: func(cmd *cobra.Command, args []string) {

			team.Configure()
		},
	}

	return cmd
}
