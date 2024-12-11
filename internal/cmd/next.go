/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/fehlhabers/zt/internal/core/handover"
	"github.com/spf13/cobra"
)

func NewNext() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "next",
		Short: "Handover to the next member of the mob/pair",
		Long:  `Hands over to the next, by committing any changes in the ztream branch and pushing to remote.`,
		Run: func(cmd *cobra.Command, args []string) {
			handover.Next()
		},
	}

	return cmd
}
