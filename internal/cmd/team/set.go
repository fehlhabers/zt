package team

import (
	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/global"
	"github.com/spf13/cobra"
)

func NewSwitch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "switch <team>",
		Short: "Switch your team",
		Long:  `Set the team which is used to identify your timer among others`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) != 1 {
				log.Fatal("Command requires 'team' to be included as argument")
			}
			err := global.GetStateKeeper().GetConfigRepo().SwitchTeam(args[0])
			if err != nil {
				log.Error("Unable to switch team", "error", err)
			}

		},
	}

	return cmd
}
