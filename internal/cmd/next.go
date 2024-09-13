/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/fehlhabers/st/internal/core/handover"
	"github.com/spf13/cobra"
)

func NewNext() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "next",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			handover.Next()
		},
	}

	return cmd
}
