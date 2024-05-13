package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Use this command to run a project",
	Long:  `Use this command to run a project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello world")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
