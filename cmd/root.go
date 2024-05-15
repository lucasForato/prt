package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Verbose bool

var rootCmd = &cobra.Command{
	Use:   "prt",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Display more verbose output in console output. (default: false)")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	var home, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

  path := filepath.Join(home, ".config", "prt")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigFile("config")
  viper.SetConfigType("toml")
  viper.AddConfigPath(".")

	viperErr := viper.ReadInConfig()
  if err != nil {
    panic(fmt.Errorf("fatal error config file: %w", viperErr))
  }

  homeDir := viper.GetString("home")

  fmt.Println(homeDir)
}
