package cmd

import (
	"github.com/discoriver/toadstool/toadstool"
	"github.com/spf13/cobra"
)

func init() {
	// Add Command
	rootCmd.AddCommand(uninstallCmd)
}

var uninstallCmd = &cobra.Command {
	Use: "uninstall",
	Short: "Uninstall a gnome extension",
	Long: "Uninstall a Gnome extension with the given name.",
	Run: func(cmd *cobra.Command, args []string) {
		toadstool.UninstallExtension()
	},
}