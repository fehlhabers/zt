/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"github.com/fehlhabers/zt/internal/core/handover"
	"github.com/spf13/cobra"
)

func NewCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new ztream",
		Long:  `Prepare a new ztream `,
		Run: func(cmd *cobra.Command, args []string) {

			var (
				ztreamName     string
				ztreamMetadata string
				err            error
			)
			if len(args) == 1 {
				ztreamName = args[0]
				ztreamMetadata = ""
			} else {
				ztreamName, ztreamMetadata, err = fromForm()
				if err != nil {
					os.Exit(1)
				}
			}

			log.Info("Creating ztream..")
			handover.CreateZtream(ztreamName, ztreamMetadata)
		},
	}
	return cmd
}

func fromForm() (ztreamName string, ztreamMetadata string, err error) {
	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Ztream name").
				Description("Also used as branch name").
				Value(&ztreamName),

			huh.NewText().
				Title("Ztream metadata").
				Description("Information which should be added to commits such as task references").
				Value(&ztreamMetadata).
				CharLimit(400).
				Lines(5),
		),
	).Run(); err != nil {
		log.Error("Error while creating ztream", "error", err)
		return "", "", err
	}

	return ztreamName, ztreamMetadata, nil
}
