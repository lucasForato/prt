package cmd

import (
	"fmt"
	"log"
	"os"
  "path/filepath"
  
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Use this command to initialize prt",
	Long:  `Use this command to initialize prt`,
	Run: func(cmd *cobra.Command, args []string) {
    var home, err = os.UserHomeDir()
    if err != nil {
      log.Fatal(err)
    }

    path := filepath.Join(home, ".config")
		err = os.Chmod(path, 0755)
		if err != nil {
			log.Fatal(err)
		}

    path = filepath.Join(home, ".config", "prt")
		err = os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Initialized prt")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
