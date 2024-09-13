/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/fehlhabers/st/internal/core/handover"
	"github.com/spf13/cobra"
)

func NewCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new ztream",
		Long:  `Prepare a new ztream `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Creating session...")
			handover.CreateSession(args[0])
		},
	}
	return cmd
}
