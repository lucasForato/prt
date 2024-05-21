package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/lucasForato/prt/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var git bool
var terms int
var cmds []string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Use this command to run a project",
	Long: `Use this command to create a tmux session for a project listed on the config.

  Examples:
  [prt run my-project -g] -> Starts a session at my-project with a lazygit window.
  [prt run my-project -t 2] -> Starts a session at my-project with two terminal windows.
  `,
	Args: cobra.ExactArgs(1),
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
			log.WithFields(log.Fields{
				"project": project,
			}).Fatal("Project not found.")
		}

		if utils.InTmuxSession() && !utils.SessionExists(project) {
			log.Fatal("Kill the current session using [kill] before initializing a new session.")
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

		if terms > 3 {
			log.WithFields(log.Fields{
				"provided": terms,
			}).Fatal("The maximum number of terminal panes is 3.")
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
	runCmd.Flags().BoolVarP(&git, "git", "g", false, "Includes a lazygit window to the session.")
	runCmd.Flags().IntVarP(&terms, "terms", "t", 2, "Includes a window with the specified number of terminal panes.")
	runCmd.Flags().Lookup("terms").NoOptDefVal = "2"
	rootCmd.AddCommand(runCmd)
}
