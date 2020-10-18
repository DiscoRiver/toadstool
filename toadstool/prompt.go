package toadstool

import (
	"github.com/manifoldco/promptui"
	"log"
	"os"
)

func promptSelect(label string, items []string) string {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Fatal("prompt failure, cannot continue.")
	}
	return result
}

func askYesNo(question string) string {
	return promptSelect(question, []string{"yes", "no"})
}

func askYesNoExit(question string) string {
	if r := promptSelect(question, []string{"yes", "no"}); r == "no" {
		os.Exit(0)
	}
	return "yes"
}