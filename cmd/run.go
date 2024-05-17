package cmd

import (
	"log"
	"os"

	"github.com/lucasForato/prt/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var git bool
var terms bool
var term bool
var cmds []string

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

		if utils.InTmuxSession() && !utils.SessionExists(project) {
			log.Fatal("Use [kill] before doing that.")
		} else if utils.InTmuxSession() && utils.SessionExists(project) {
			utils.SwitchSession(project)
			return
		} else if utils.SessionExists(project) {
			utils.AttachSession(project)
			return
		}

		if err := os.Chdir(dir); err != nil {
			log.Fatal(err)
		}

    numTerms := 0
    if (terms) {
      numTerms = 2
    } else if (term) {
      numTerms = 1
    }

		var tmux = utils.Tmux{
			Name:  project,
			Git:   git,
			Terms: numTerms,
      Cmd: cmds,
		}

		tmux.CreateSession()
	},
}

func init() {
	runCmd.Flags().BoolVarP(&git, "git", "g", false, "Run with this to a git window")
	runCmd.Flags().BoolVarP(&terms, "terms", "t", false, "Run with this flag to get two terminals")
	runCmd.Flags().BoolVarP(&term, "term", "e", false, "Run with this flag to get a single terminal")
  runCmd.Flags().StringSliceVarP(&cmds, "cmd", "c", cmds, "Run with this flag to add commands to the terminals")
	rootCmd.AddCommand(runCmd)
}
