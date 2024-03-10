package connect

import (
	"fmt"

	"github.com/aideneyre/acli/internal/config"
	connectConfig "github.com/aideneyre/acli/internal/config/configurator/connect"
	connectssh "github.com/aideneyre/acli/internal/connect"

	"github.com/spf13/cobra"
)

var (
	ConnectCmd = &cobra.Command{
		Use:   "connect [flags] [alias]",
		Short: "Connect to servers using SSH.",
		Long:  "Connect to servers using SSH with predefined aliases. They must be configured through the config command.",
		Example: `
# Connect to a server using an interactive prompt
acli connect

# Connect to a server using a predefined alias
acli config <alias>
		`,
		Version:      "0.1.0",
		RunE:         runE,
		SilenceUsage: true,
	}
)

// runE connects to a server using SSH.
func runE(cmd *cobra.Command, args []string) error {
	err := config.Initialize()
	if err != nil {
		return fmt.Errorf("unable to initialize acli: %w", err)
	}

	var globalConfig connectConfig.Config
	err = config.UnmarshalConfig(&globalConfig)
	if err != nil {
		return fmt.Errorf("unable to load config: %w", err)
	}

	if len(args) == 1 {
		return connectssh.Connect(globalConfig, args[0])
	}

	err = connectssh.ConnectInteractive(globalConfig)
	if err != nil {
		return fmt.Errorf("unable to connect to server: %w", err)
	}

	return nil
}
