package cmd

import (
	"github.com/discoriver/toadstool/toadstool"
	"github.com/spf13/cobra"
)

var extensionZip string

func init() {
	// Flags
	installCmd.PersistentFlags().StringVarP(&extensionZip, "extension", "e","", "Gnome extension to install")

	// Add Command
	rootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command {
	Use: "install",
	Short: "Install a Gnome extension",
	Long: "Install a Gnome extension at configured path.",
	Run: func(cmd *cobra.Command, args []string) {
		toadstool.InstallExtension(extensionZip)
	},
}


