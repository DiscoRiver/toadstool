package cmd

import (
	"github.com/discoriver/toadstool/toadstool"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List installed Gnome extensions",
	Long: "List installed Gnome extensions in the configured path",
	Run: func(cmd *cobra.Command, args []string) {
		toadstool.ListInstalledUserExtensions()
	},
}
