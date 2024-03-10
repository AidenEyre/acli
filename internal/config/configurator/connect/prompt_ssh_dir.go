package connectConfig

import (
	"fmt"

	"github.com/aideneyre/acli/internal/config"
	"github.com/aideneyre/acli/internal/prompt"
)

func promptSetSSHDirectory() error {
	directory, err := prompt.PromptForInput("Enter the path to your SSH directory")
	if err != nil {
		return fmt.Errorf("unable to prompt for SSH directory: %w", err)
	}

	globalConfig.Connect.SSHDir = directory

	err = config.Set("connect", globalConfig.Connect)
	if err != nil {
		return fmt.Errorf("error setting SSH directory: %w", err)
	}

	fmt.Println("SSH directory set successfully.")
	return nil
}
