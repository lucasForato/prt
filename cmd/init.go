package cmd

import (
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/lucasForato/prt/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize prt by creating a config file at .config/prt",
	Long: `Initialization command:

  This command will create a config file at .config/prt/
  You should run this command before using any other command. 
  `,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Initializing prt...")

		dotConfig := utils.GetDirFromHome(".config")
		err := os.Chmod(dotConfig, 0755)
		if err != nil {
			log.WithFields(log.Fields{
				"directory": ".config",
			}).Fatal("Error changing permissions of directory")
		}

		prt := utils.GetDirFromHome(".config", "prt")
		err = os.MkdirAll(prt, 0755)
		if err != nil {
			log.WithFields(log.Fields{
				"directory": ".config/prt",
			}).Fatal("Error creating directory")
		}

		config := utils.GetDirFromHome(".config", "prt", "config.yaml")
		if !utils.DirExists(config) {
			_, createErr := os.Create(config)
			if createErr != nil {
				log.WithFields(log.Fields{
					"directory": ".config/prt",
					"config":    "config.yaml",
				}).Fatal("Error creating config file")
			}
		} else {
			log.WithFields(log.Fields{
				"directory": ".config/prt",
				"file":      "config.yaml",
			}).Warn("Config file already exists")
      return
		}

		log.WithFields(log.Fields{
			"directory": ".config/prt",
			"file":      "config.yaml",
		}).Info("Config file created sucessfully")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
