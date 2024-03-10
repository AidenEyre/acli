package kchconfig

import (
	"fmt"

	"github.com/aideneyre/acli/internal/config"
	"github.com/aideneyre/acli/internal/prompt"
	"github.com/manifoldco/promptui"
)

var (
	configOptions = []string{
		"Always open K9s",
		"Set default resource type",
		"Set default namespace",
		"exit",
	}
)

func Configure() error {
	exit := false
	var err error
	for !exit {
		exit, err = runPrompt()
	}
	if err != nil {
		return err
	}
	return nil
}

func runPrompt() (bool, error) {
	option, err := promptOptions()
	if err != nil {
		return false, fmt.Errorf("unable to prompt for options: %w", err)
	}

	switch option {
	case "Always open K9s":
		err := promptAlwaysOpenK9s()
		return false, err
	case "Set default resource type":
		err := promptDefaultResourceType()
		return false, err
	case "Set default namespace":
		err := promptDefaultNamespace()
		return false, err
	case "exit":
		return true, nil
	default:
		return false, nil
	}
}

func promptOptions() (string, error) {
	options := promptui.Select{
		Label:        "Select an option to configure",
		Items:        configOptions,
		Stdout:       &prompt.BellSkipper{},
		Size:         15,
		HideHelp:     true,
		HideSelected: true,
	}

	_, option, err := options.Run()
	if err != nil {
		return "", err
	}
	return option, nil
}

func promptAlwaysOpenK9s() error {
	prompt := promptui.Select{
		Label:        "Always open K9s",
		Items:        []string{"Yes", "No"},
		Stdout:       &prompt.BellSkipper{},
		Size:         15,
		HideHelp:     true,
		HideSelected: true,
	}

	_, choice, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("unable to prompt for always open K9s: %w", err)
	}

	if choice == "Yes" {
		config.Set("kch.alwaysOpenK9s", true)
	} else {
		config.Set("kch.alwaysOpenK9s", false)
	}

	return nil
}

func promptDefaultResourceType() error {
	prompt := promptui.Prompt{
		Label: "Set default resource type",
	}

	choice, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("unable to prompt for default resource type: %w", err)
	}

	config.Set("kch.defaults.Resource", choice)
	return nil
}

func promptDefaultNamespace() error {
	prompt := promptui.Prompt{
		Label: "Set default namespace",
	}

	choice, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("unable to prompt for default namespace: %w", err)
	}

	config.Set("kch.defaults.namespace", choice)
	return nil
}
