package cmd

import (
	"github.com/lucasForato/prt/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Use this command to delete a project from the configuration",
	Long:  `This command will remove and entry from the configuration file at /home/.config/prt/config.yaml`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		project := args[0]

		config := utils.GetConfigDir()
		file := utils.GetConfigFile()
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(config)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}

		configMap := viper.AllSettings()
		for key := range configMap {
			if key == project {
				delete(configMap, key)
			}
		}

		var entries []string

		for key, val := range configMap {
			if strVal, ok := val.(string); ok {
				entries = append(entries, utils.BuildEntry(key, strVal))
			} else {
				log.Fatal("Config contains an incorrect value")
			}
		}

		utils.RewriteFile(file, entries)

		log.WithFields(log.Fields{
			"project": project,
		}).Info("Project deleted successfully from configuration file.")
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
