package cliconfigsetup

import (
	"errors"
	"fmt"
	"os"

	"github.com/aideneyre/acli/internal/config"
	"github.com/aideneyre/acli/internal/config/configurator"
	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

var (
	ConfigCmd = &cobra.Command{
		Use:       "config [flags] [options]",
		Short:     "Configure acli settings.",
		Long:      "Opens an interactive prompt to configure acli settings for select subcommands: kch, connect.",
		ValidArgs: []string{"view"},
		Example: `
# Open the configuration prompt
acli config

# View the existing configuration
acli config view
		`,
		Version:      "0.1.0",
		RunE:         runE,
		SilenceUsage: true,
	}
)

// runE calls the setup function in the config package.
func runE(cmd *cobra.Command, args []string) error {
	err := config.Initialize()
	if err != nil {
		return fmt.Errorf("unable to initialize acli: %w", err)
	}

	if len(args) == 1 && args[0] == "view" {
		return printConfig()
	}

	err = configurator.RunSetup()
	if errors.Is(err, promptui.ErrInterrupt) {
		fmt.Print("\033[u\033[J") // Clear the terminal
		fmt.Println("CLI exited early!")
		os.Exit(0)
	}
	if err != nil {
		return fmt.Errorf("unable to configure acli: %w", err)
	}

	return nil
}

func printConfig() error {
	var yamlConfig interface{}
	err := config.UnmarshalConfig(&yamlConfig)
	if err != nil {
		return fmt.Errorf("unable to unmarshal config: %w", err)
	}

	yamlBytes, err := yaml.Marshal(&yamlConfig)
	if err != nil {
		return fmt.Errorf("unable to marshal config to YAML: %w", err)
	}

	fmt.Println(string(yamlBytes))
	return nil
}
