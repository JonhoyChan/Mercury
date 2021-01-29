package cmd

import (
	"github.com/spf13/cobra"
)

func NewMercuryCommand(repoPath string) (*cobra.Command, func() <-chan error) {
	opt := NewOptions(repoPath)

	cmd := &cobra.Command{
		Use:   "mercury",
		Short: "",
		Long:  ``,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
	}

	cmd.PersistentFlags().StringVar(&opt.repoPath, "repo", repoPath, "filepath to load mercury data from")
	cmd.PersistentFlags().StringVar(&opt.infraUrl, "infra", "http://localhost:9600/infra/v1", "infra url to load remote config from")

	cmd.AddCommand(
		NewInfraCommand(opt),
		NewJobCommand(opt),
		NewCometCommand(opt),
		NewLogicCommand(opt),
	)
	return cmd, opt.Shutdown
}
