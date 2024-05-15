package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Verbose bool

var rootCmd = &cobra.Command{
	Use:   "prt",
	Short: "short description.",
	Long:  `long description.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Display more verbose output in console output. (default: false)")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}
