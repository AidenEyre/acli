package cliconfigsetup

import (
	"errors"
	"fmt"
	"os"

	"github.com/aideneyre/acli/internal/config"
	"github.com/aideneyre/acli/internal/config/configurator"
	"github.com/manifoldco/promptui"

	"github.com/spf13/cobra"
)

var (
	ConfigCmd = &cobra.Command{
		Use:   "config [flags]",
		Short: "Configure acli settings.",
		Long:  "Opens an interactive prompt to configure acli settings for select subcommands: kch, connect.",
		Example: `
# Open the configuration prompt
acli config
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
