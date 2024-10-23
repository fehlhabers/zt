package team

import (
	"github.com/spf13/cobra"
)

func NewTeam() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "team",
		Short: "Handle different team commands",
		Long:  `Team commands such as setting, listing teams, etc`,
	}

	cmd.AddCommand(NewSet())
	return cmd
}
