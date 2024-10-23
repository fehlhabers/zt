/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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
			log.Info("Creating ztream...")
			handover.CreateZtream(args[0])
		},
	}
	return cmd
}
