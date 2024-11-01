package connectConfig

import (
	"fmt"

	"github.com/aideneyre/acli/internal/config"
	"github.com/aideneyre/acli/internal/prompt"
	"github.com/manifoldco/promptui"
)

func promptModifyDeleteSSHAlias() error {
	if len(globalConfig.Connect.Aliases) == 0 {
		fmt.Println("No existing SSH aliases to modify or delete.")
		return nil
	}

	var aliasNames []string
	for _, alias := range globalConfig.Connect.Aliases {
		aliasNames = append(aliasNames, alias.Name)
	}

	selectPrompt := promptui.Select{
		Label:  "Select an alias to modify or delete",
		Items:  aliasNames,
		Stdout: &prompt.BellSkipper{},
	}
	aliasInex, _, err := selectPrompt.Run()
	if err != nil {
		return fmt.Errorf("selecting alias failed: %w", err)
	}
	selectedAlias := globalConfig.Connect.Aliases[aliasInex]
	fmt.Printf("Selected alias: %s\n", selectedAlias.Name)

	actionPrompt := promptui.Select{
		Label:  "Would you like to modify or delete this alias?",
		Items:  []string{"Modify", "Delete", "Cancel"},
		Stdout: &prompt.BellSkipper{},
	}
	_, action, err := actionPrompt.Run()
	if err != nil {
		return fmt.Errorf("selecting action failed: %w", err)
	}

	switch action {
	case "Modify":
		if err := modifyAlias(aliasInex); err != nil {
			return err
		}
	case "Delete":
		globalConfig.Connect.Aliases = append(globalConfig.Connect.Aliases[:aliasInex], globalConfig.Connect.Aliases[aliasInex+1:]...)
		err = config.Set("connect", globalConfig.Connect)
		if err != nil {
			return err
		}
		fmt.Println("SSH alias deleted successfully.")
	case "Cancel":
		return nil
	}

	return nil
}

func modifyAlias(index int) error {
	alias := &globalConfig.Connect.Aliases[index]

	newName, err := prompt.PromptForInputWithDefault("Enter new alias name (leave blank to keep current)", alias.Name)
	if err != nil {
		return err
	}
	newPEMFile, err := prompt.PromptForInputWithDefault("Enter new PEM file name (leave blank to keep current)", alias.PEMFile)
	if err != nil {
		return err
	}
	newIP, err := prompt.PromptForInputWithDefault("Enter new IP address (leave blank to keep current)", alias.IP)
	if err != nil {
		return err
	}
	newPort, err := prompt.PromptForInputWithDefault("Enter new port number (leave blank to keep current)", alias.Port)
	if err != nil {
		return err
	}
	newUser, err := prompt.PromptForInputWithDefault("Enter new username (leave blank to keep current)", alias.User)
	if err != nil {
		return err
	}

	if newName != "" {
		alias.Name = newName
	}
	if newPEMFile != "" {
		alias.PEMFile = newPEMFile
	}
	if newIP != "" {
		alias.IP = newIP
	}
	if newPort != "" {
		alias.Port = newPort
	}
	if newUser != "" {
		alias.User = newUser
	}

	err = config.Set("connect", globalConfig.Connect)
	if err != nil {
		return err
	}
	fmt.Println("SSH alias modified successfully.")
	return nil
}
