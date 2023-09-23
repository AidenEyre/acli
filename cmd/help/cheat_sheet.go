package cmd

import (
	"fmt"

	"github.com/aideneyre/acli/internal/common"
	"github.com/aideneyre/acli/internal/help"

	"github.com/spf13/cobra"
)

var CheatSheet = &cobra.Command{
	Use:   "idk",
	Short: "Print out references for common tasks and commands",
	Example: `
# Call specific cheat sheet
acli idk <page>

# Call cheat sheet with an interactive promp
acli idk
		`,
	Version:      "1.0.0",
	RunE:         runE,
	SilenceUsage: true,
}

// runE calls the GetReferencePage function and prints the returned help page
func runE(cmd *cobra.Command, args []string) error {
	var helpText string
	var err error

	if len(args) == 1 {
		helpText, err = help.GetReferencePage(args[0])
	} else {
		helpText, err = help.GetReferencePage("")
	}
	if err != nil {
		return fmt.Errorf("failed to get help text, %w", err)
	}

	common.PrintStringWithMore(helpText)
	return nil
}
