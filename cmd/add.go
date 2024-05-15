package cmd

import (
	"fmt"

	"github.com/lucasForato/prt/utils"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Use this command to add a project to config",
	Long:  `Use this command to add a project to config`,
  Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
    pathFromHome := utils.GetDirFromHome(".config", "prt", "config.yaml")
    pathFromCurrent := utils.GetDirFromCurr(args[1])

    entry := buildEntry(args[0], pathFromCurrent)
    utils.AppendToFile(pathFromHome, entry)
	},
}

func buildEntry(projectName string, projectPath string) string {
    return fmt.Sprintf("%v: \"%v\"\n", projectName, projectPath)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
