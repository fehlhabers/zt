package cmd

import (
	"github.com/fehlhabers/zt/internal/core/handover"
	"github.com/spf13/cobra"
)

func NewStart() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "start",
		Short: "Starts a new session in the active ztream",
		Long:  `Used when the next member is about to start within a mob/pair after having joined/created a ztream`,
		Run: func(cmd *cobra.Command, args []string) {
			handover.Start()
		},
	}
	return cmd
}
