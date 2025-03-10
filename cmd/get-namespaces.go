package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var getNamespacesCmd = &cobra.Command{
	Use:     "gns",
	Aliases: []string{"get-ns", "get-namespaces"},
	Short:   "List namespaces",
	Long:    "List namespaces in current kube context",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Listing namesapces:\n%s", GetNamespaces())
	},
}

func init() {
	rootCmd.AddCommand(getNamespacesCmd)
}
