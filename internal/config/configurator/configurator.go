package configurator

import (
	"fmt"

	kchconfig "github.com/aideneyre/acli/internal/config/configurator/kch"

	"github.com/aideneyre/acli/internal/prompt"
	"github.com/manifoldco/promptui"
)

var (
	subcommands = []string{
		"kch",
	}
)

func RunSetup() error {
	subcommand, err := promptSubcommands()
	if err != nil {
		return fmt.Errorf("unable to prompt for subcommand: %w", err)
	}

	switch subcommand {
	case "kch":
		err := kchconfig.Configure()
		return err
	default:
		return fmt.Errorf("subcommand %s not found", subcommand)
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
