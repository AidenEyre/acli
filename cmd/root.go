// Package cmd provides commands for the acli CLI application.
package cmd

import (
	"fmt"
	"os"

	cliconfigsetup "github.com/aideneyre/acli/cmd/config"
	"github.com/aideneyre/acli/cmd/connect"
	help "github.com/aideneyre/acli/cmd/help"
	kube "github.com/aideneyre/acli/cmd/kube"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "acli [flags] [options]",
	Short:   "Aiden CLI - quality of life commands",
	Version: "1.3.0",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("you are using acli version %s\n\n", cmd.Version)
		cmd.Help()
	},
}

// init adds all the commands to the root command, allowing them to be executed
// via the CLI
func init() {
	RootCmd.AddCommand(kube.KubeChangeCmd)
	RootCmd.AddCommand(help.CheatSheet)
	RootCmd.AddCommand(cliconfigsetup.ConfigCmd)
	RootCmd.AddCommand(connect.ConnectCmd)
}

// Execute runs the command and its subcommands, returning an error
func Execute() {
	fmt.Print("\033[s") // Save cursor position so commands can clear terminal text

	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
