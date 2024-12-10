/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/fehlhabers/zt/internal/core/handover"
	"github.com/spf13/cobra"
)

func NewMerge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "merge",
		Short: "Merge the ztream into main using the teams merge strategy",
		Long:  `Merge the ztream into main using the teams merge strategy`,
		Run: func(cmd *cobra.Command, args []string) {
			handover.Merge()
		},
	}

	return cmd
}
