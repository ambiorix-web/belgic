package cmd

import (
	"log"

	"github.com/devOpifex/eburon/internal/config"
	"github.com/spf13/cobra"
)

var conf string

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Create configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		err := config.CheckConfigPath(conf)
		if err != nil {
			log.Fatal(err)
		}
		err = config.Create(conf)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	configCmd.Flags().StringVarP(&conf, "path", "p", "/eburon", "Path to the config file")
	rootCmd.AddCommand(configCmd)
}
