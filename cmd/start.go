package cmd

import (
	"github.com/devOpifex/belgic/internal"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Run()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
