package team

import (
	"errors"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/core/config"
	"github.com/fehlhabers/zt/internal/domain"
	"github.com/spf13/cobra"
)

func NewConfigure() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Configure a team",
		Long:  `Configure a new team with setting using a wizard`,
		Run: func(cmd *cobra.Command, args []string) {

			var (
				name           string
				sessionMinsStr string
				mainBranch     string
				confirm        bool
			)

			if err := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("Enter team name").
						Value(&name).
						// Validating fields is easy. The form will mark erroneous fields
						// and display error messages accordingly.
						Validate(func(str string) error {
							if str == "Frank" {
								return errors.New("Sorry, we don’t serve customers named Frank.")
							}
							return nil
						}),

					huh.NewInput().
						Title("What's your default session length?").
						Description("In minutes").
						Placeholder("10").
						Value(&sessionMinsStr).
						Validate(func(str string) error {
							mins, err := strconv.Atoi(str)
							if err != nil {
								return errors.New("must be valid integer (0-60)")
							}

							if mins < 1 || mins > 60 {
								return errors.New("must be valid integer (0-60)")
							}
							return nil
						}),

					huh.NewInput().
						Title("What branch is used as trunk/main/master?").
						Value(&mainBranch),

					huh.NewConfirm().
						Title("Add Team?").
						Value(&confirm),
				),
			).Run(); err != nil {
				log.Error("Failed to conclude configuration")
			}

			if !confirm {
				log.Fatal("Configuration aborted")
			}

			sessionMins, _ := strconv.Atoi(sessionMinsStr)

			teamCfg := &domain.TeamConfig{
				SessionDurMins: sessionMins,
				MainBranch:     mainBranch,
			}
			config.AddTeam(name, *teamCfg)
			log.Infof("Team %s added!", name)
		},
	}

	return cmd
}
