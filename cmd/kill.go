package cmd

import (
	"github.com/lucasForato/prt/utils"
	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Use this command to kill the current session",
	Long:  `Use this command to kill the current session`,
  Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		utils.KillCurrSession()
	},
}

func init() {
	rootCmd.AddCommand(killCmd)
}
