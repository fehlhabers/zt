/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "st",
		Short: "Create Ztreams in your team!",
		Long: `Create quick git handovers, by setting up ztreams in your team.
Create, join, switch to next and close ztreams. Your team can work in multiple parallel ztreams
by leveraging different git branches`,
	}

	cmd.AddCommand(NewStart())
	cmd.AddCommand(NewNext())
	cmd.AddCommand(NewCreate())

	return cmd
}
