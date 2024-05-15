package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Use this command to add a project to config",
	Long:  `Use this command to add a project to config`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("add command")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
