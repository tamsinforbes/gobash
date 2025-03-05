package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"listfiles"},
	Short:   "List files",
	Long:    "List files in current directory",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Listing files:\n%s", List(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
