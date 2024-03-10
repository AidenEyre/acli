package connectConfig

import (
	"fmt"

	"github.com/aideneyre/acli/internal/config"
	"github.com/aideneyre/acli/internal/prompt"
	"github.com/manifoldco/promptui"
)

// Config represents the global configuration structure
type Config struct {
	Connect ConnectConfig `yaml:"connect"`
}

// ConnectConfig represents the SSH connection configurations
type ConnectConfig struct {
	Aliases []SSHAlias `yaml:"aliases"`
	SSHDir  string     `yaml:"sshdir"`
}

// SSHAlias represents a single SSH alias configuration
type SSHAlias struct {
	Name    string `yaml:"name"`
	PEMFile string `yaml:"pemfile"`
	IP      string `yaml:"ip"`
	User    string `yaml:"user"`
}

var (
	configOptions = []string{
		"Add new SSH alias",
		"Modify/Delete SSH aliases",
		"Set SSH directory",
		"exit",
	}
	globalConfig Config
)

func loadConfig() error {
	if err := config.UnmarshalConfig(&globalConfig); err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}
	return nil
}

func Configure() error {
	exit := false
	var err error

	err = loadConfig()
	if err != nil {
		return fmt.Errorf("unable to load config: %w", err)
	}

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
	case "Add new SSH alias":
		err := promptAddNewSSHAlias()
		return false, err
	case "Modify/Delete SSH aliases":
		err := promptModifyDeleteSSHAlias()
		return false, err
	case "Set SSH directory":
		err := promptSetSSHDirectory()
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
