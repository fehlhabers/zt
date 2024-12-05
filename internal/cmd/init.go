/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/fehlhabers/zt/internal/core/team"
	"github.com/spf13/cobra"
)

func NewInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initial config",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			team.Configure()
		},
	}

	return cmd
}
