package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/lucasForato/prt/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Use this command to initialize prt",
	Long:  `Use this command to initialize prt`,
	Run: func(cmd *cobra.Command, args []string) {
		dotConfig := utils.GetDirFromHome(".config")
		err := os.Chmod(dotConfig, 0755)
		if err != nil {
			log.Fatal(err)
		}

		prt := utils.GetDirFromHome(".config", "prt")
		err = os.MkdirAll(prt, 0755)
		if err != nil {
			log.Fatal(err)
		}

		config := utils.GetDirFromHome(".config", "prt", "config.yaml")
		if !utils.DirExists(config) {
			_, createErr := os.Create(config)
			if createErr != nil {
				log.Fatal(createErr)
			}
		} else {
      fmt.Println("prt already initialized.")
    }

		if Verbose {
			fmt.Println("Initialized prt.")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
