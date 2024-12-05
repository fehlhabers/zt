/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/fehlhabers/zt/internal/cmd/team"
	"github.com/spf13/cobra"
)

func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "zt",
		Short: "Create Ztreams in your team!",
		Long: `Create quick git handovers, by setting up ztreams in your team.
Create, join, switch to next and close ztreams. Your team can work in multiple parallel ztreams
by leveraging different git branches`,
	}

	cmd.AddCommand(NewStart())
	cmd.AddCommand(NewJoin())
	cmd.AddCommand(NewNext())
	cmd.AddCommand(NewCreate())
	cmd.AddCommand(team.NewTeam())

	return cmd
}
