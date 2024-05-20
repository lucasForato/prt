package cmd

import (
	"github.com/lucasForato/prt/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Use this command to add a project to config",
	Long: `This commands adds a new project at /home/.config/prt/config.yaml

  - Use [prt ls] to list all projects at the configuration file.
  `,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		project := args[0]

		config := utils.GetConfigDir()
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(config)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}

		if exists := viper.Get(project); exists != nil {
			log.WithFields(log.Fields{
				"project": project,
			}).Fatal("This project already exists.")
		}

		dir := utils.GetDirFromCurr(args[1])
		entry := utils.BuildEntry(project, dir)

    config = utils.GetConfigFile()
		utils.AppendToFile(config, entry)
		log.WithFields(log.Fields{
			"project":   project,
			"directory": dir,
		}).Info("Added new project to prt")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
