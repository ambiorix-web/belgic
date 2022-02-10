package cmd

import (
	"fmt"
	"os"

	"github.com/devOpifex/eburon/internal"
	"github.com/spf13/cobra"
)

var conf string

var rootCmd = &cobra.Command{
	Use:   "eburon",
	Short: "Eburon is a web server for {ambiorix} applications",
	Long: `A webserver to easily manage {ambiorix} applications
  	which also eases the management of concurrent users.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.Run(conf)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&conf, "config", "/eburon/eburon.config", "location of the config file")
}
