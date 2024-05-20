package cmd

import (
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/lucasForato/prt/utils"
	"github.com/spf13/cobra"
)

var Verbose bool

var rootCmd = &cobra.Command{
	Use:   "prt",
	Short: "Prt let's you initialize tmux session effortlessly by using simple commands.",
	Long: `Prt is your advanced tmux session manager

  By using the prt command, you can initialize tmux session effortlessly by only providing command line arguments.

  - use [prt] followed by a directory to initialize the default session.
  - use [prt init] to initialize a configuration.
  `,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]

		if utils.InTmuxSession() {
			log.Fatal("Kill the current session using [kill] before initializing a new session.")
		}

		if err := os.Chdir(dir); err != nil {
			log.Fatal(err)
		}

		var tmux = utils.Tmux{
			Name:  "prt",
			Git:   false,
			Terms: 1,
		}

		tmux.CreateSession()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
