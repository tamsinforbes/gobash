package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gobash",
	Short: "gobash is a cli tool for wrapping bash commands",
	Long:  "gobash is a cli tool for wrapping bash commands - such as? who knows?",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Gobash '%s'\n", err)
		os.Exit(1)
	}
}
