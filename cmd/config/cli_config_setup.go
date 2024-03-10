package cliconfigsetup

import (
	"fmt"

	"github.com/aideneyre/acli/internal/config/configurator"

	"github.com/spf13/cobra"
)

var (
	ConfigCmd = &cobra.Command{
		Use:   "config [flags]",
		Short: "Configure acli settings.",
		Long:  "Opens an interactive prompt to configure acli settings for select subcommands: kch.",
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
	err := configurator.RunSetup()
	if err != nil {
		return fmt.Errorf("unable to configure acli: %w", err)
	}

	return nil
}
