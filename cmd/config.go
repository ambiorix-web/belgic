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
	Short: "Create a configuration file",
	Long:  `Specify a directory in which to create the config file, see -p flag to specify path`,
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
	configCmd.Flags().StringVarP(&conf, "path", "p", "", "Path to the config file")
	rootCmd.AddCommand(configCmd)
}
