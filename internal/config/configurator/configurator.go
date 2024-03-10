package configurator

import (
	"fmt"

	connectConfig "github.com/aideneyre/acli/internal/config/configurator/connect"
	kchconfig "github.com/aideneyre/acli/internal/config/configurator/kch"

	"github.com/aideneyre/acli/internal/prompt"
	"github.com/manifoldco/promptui"
)

var (
	subcommands = []string{
		"kch",
		"connect",
		"exit",
	}
)

func RunSetup() error {
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
	subcommand, err := promptSubcommands()
	if err != nil {
		return true, fmt.Errorf("unable to prompt for subcommand: %w", err)
	}

	switch subcommand {
	case "kch":
		err := kchconfig.Configure()
		return false, err
	case "connect":
		err := connectConfig.Configure()
		return false, err
	case "exit":
		return true, nil
	default:
		return true, fmt.Errorf("subcommand %s not found", subcommand)
	}
}

func promptSubcommands() (string, error) {
	promptOptions := promptui.Select{
		Label:        "Select a subcommand to configure",
		Items:        subcommands,
		Stdout:       &prompt.BellSkipper{},
		Size:         15,
		HideHelp:     true,
		HideSelected: true,
	}

	_, subcommand, err := promptOptions.Run()
	if err != nil {
		return "", err
	}

	return subcommand, nil
}
