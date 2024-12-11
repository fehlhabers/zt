/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/fehlhabers/zt/internal/core/admin"
	"github.com/spf13/cobra"
)

func NewList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists all ztreams",
		Long:  `Lists all the ztreams. Ztreams that are not active are mostly for reference purposes at this time.`,
		Run: func(cmd *cobra.Command, args []string) {
			admin.ListZtreams()
		},
	}

	return cmd
}
