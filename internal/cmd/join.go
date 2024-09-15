package cmd

import (
	"github.com/fehlhabers/zt/internal/core/handover"
	"github.com/spf13/cobra"
)

func NewJoin() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "join",
		Short: "Join an existing ztream",
		Long:  `If a team mate has already created a ztream, use this command to join it. A ztream usually corresponds to a git branch.`,
		Run: func(cmd *cobra.Command, args []string) {
			handover.JoinZtream(args[0])
		},
	}
	return cmd
}
