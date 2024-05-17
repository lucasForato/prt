package cmd

import (
	"log"
	"os"

	"github.com/lucasForato/prt/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var git bool
var terms int

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Use this command to run a project",
	Long:  `Use this command to run a project`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		project := args[0]

		config := utils.GetConfigDir()
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(config)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}

		dir := viper.GetString(project)
		if len(dir) == 0 {
			log.Fatal("Project not found.")
		}

		if !utils.SessionExists(project) && utils.InTmuxSession() {
			log.Fatal("Use [kill] before doing that.")
		} else if utils.SessionExists(project) && utils.InTmuxSession() {
			utils.SwitchSession(project)
			return
		} else if utils.SessionExists(project) {
			utils.AttachSession(project)
			return
		}

		if err := os.Chdir(dir); err != nil {
			log.Fatal(err)
		}

    if terms > 2 {
      log.Fatal("You can't use more than 2 terminals.")
    }

		var tmux = utils.Tmux{
			Name:  project,
			Git:   git,
			Terms: terms,
		}

		tmux.CreateSession()
	},
}

func init() {
	runCmd.Flags().BoolVarP(&git, "git", "g", false, "Run with this to a git window")
	runCmd.Flags().IntVarP(&terms, "term", "t", terms, "Number of terminals")
	rootCmd.AddCommand(runCmd)
}
