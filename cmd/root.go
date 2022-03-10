package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "belgic",
	Short: "Belgic is a web server for {ambiorix} applications",
	Long:  `A load balancer to easily scale an {ambiorix} application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// do nothing
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
