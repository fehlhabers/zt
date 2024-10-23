package team

import (
	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/adapter/state"
	"github.com/fehlhabers/zt/internal/model"
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

			zt := model.Team{Name: args[0]}
			state.Storer.CreateTeam(zt)
		},
	}

	return cmd
}
