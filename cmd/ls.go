package cmd

import (
	"fmt"

	"github.com/lucasForato/prt/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Use this command to list all available projects",
	Long:  `This command lists all projects at .config/prt/config.yaml by providing the project name and directory where the project is nested`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.GetConfigDir()
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(config)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}

		values := viper.AllSettings()
		biggerThanZero := len(values) > 0

		if biggerThanZero {
			fmt.Println()
		}
		for i, v := range values {
			log.Info("[", i, "]: ", v)
		}
		if biggerThanZero {
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
