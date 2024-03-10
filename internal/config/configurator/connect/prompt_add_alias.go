package connectConfig

import (
	"fmt"

	"github.com/aideneyre/acli/internal/config"
	"github.com/aideneyre/acli/internal/prompt"
)

func promptAddNewSSHAlias() error {
	var aliasName, pemFileName, ipAddress, userName string
	var err error

	aliasName, err = prompt.PromptForInput("Enter alias name")
	if err != nil {
		return fmt.Errorf("unable to prompt for alias name: %w", err)
	}
	pemFileName, err = prompt.PromptForInput("Enter PEM file name")
	if err != nil {
		return fmt.Errorf("unable to prompt for PEM file name: %w", err)
	}
	ipAddress, err = prompt.PromptForInput("Enter IP address")
	if err != nil {
		return fmt.Errorf("unable to prompt for IP address: %w", err)
	}
	userName, err = prompt.PromptForInput("Enter username:")
	if err != nil {
		return fmt.Errorf("unable to prompt for username: %w", err)
	}

	newAlias := SSHAlias{
		Name:    aliasName,
		PEMFile: pemFileName,
		IP:      ipAddress,
		User:    userName,
	}

	globalConfig.Connect.Aliases = append(globalConfig.Connect.Aliases, newAlias)

	err = config.Set("connect", globalConfig.Connect)
	if err != nil {
		fmt.Println("Error setting new SSH alias:", err)
		return err
	}

	fmt.Println("New SSH alias added successfully.")
	return nil
}
