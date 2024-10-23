package cmd

import (
	"github.com/fehlhabers/zt/internal/assert"
	"github.com/fehlhabers/zt/internal/core/handover"
	"github.com/spf13/cobra"
)

func NewStart() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "start",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			assert.Start()
			handover.Start()
		},
	}
	return cmd
}
