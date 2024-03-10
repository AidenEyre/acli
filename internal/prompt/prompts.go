package prompt

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func PromptForInput(label string) (string, error) {
	prompt := promptui.Prompt{
		Label:  label,
		Stdout: &BellSkipper{},
	}

	result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %w", err)
	}

	return result, nil
}

func PromptForInputWithDefault(label, defaultValue string) (string, error) {
	prompt := promptui.Prompt{
		Label:     label,
		Default:   defaultValue,
		Stdout:    &BellSkipper{},
		AllowEdit: true,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %w", err)
	}

	return result, nil
}
