package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/lucasForato/prt/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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

		if utils.SessionExists(project) {
			utils.KillSession(project)
		}

		if utils.InTmuxSession() {
			utils.KillCurrSession()
		}

		if err := os.Chdir(dir); err != nil {
			log.Fatal(err)
		}

		tmuxCmd := exec.Command("tmux", "new-session", "-s", project)
		tmuxCmd.Stdin = os.Stdin
		tmuxCmd.Stdout = os.Stdout
		tmuxCmd.Stderr = os.Stderr

		fmt.Println("Session started.")
		if err := tmuxCmd.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
