package toadstool

import "fmt"

func UninstallExtension() {
	installedExtensions := getInstalledUserExtensions()

	extensionToUninstall := promptSelect("Please select an extension to uninstall:", installedExtensions)

	err := removeDirectory(fmt.Sprintf("%s%s%s", extensionsDirectory, "/", extensionToUninstall))
	if err != nil {
		fmt.Printf("Couldn't remove extension %s: %s", extensionToUninstall, err)
		return
	}

	restartGnomeShell()
}