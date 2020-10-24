package toadstool

import (
	"fmt"
	"os/exec"
)

func restartGnomeShell() {
	cmd := exec.Command("/bin/bash", "-c", "busctl --user call org.gnome.Shell /org/gnome/Shell org.gnome.Shell Eval s 'Meta.restart(\"Restarting…\")'")

	fmt.Println("Restarting Gnome Shell...")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Unable to restart Gnome shell: ", err)
	}
	fmt.Println("Uninstall complete!")
}


