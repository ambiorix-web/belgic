package cmd

import (
	"log"

	"github.com/devOpifex/belgic/internal/config"
	"github.com/spf13/cobra"
)

// conf path to the configuration file to create.
var conf string

// configCmd is a command to create an initial config file.
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
	configCmd.Flags().StringVarP(&conf, "path", "p", "/belgic", "Path to the config file")
	rootCmd.AddCommand(configCmd)
}
