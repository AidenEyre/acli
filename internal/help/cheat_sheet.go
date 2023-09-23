package help

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/aideneyre/acli/internal/help/helprefs"
	"github.com/aideneyre/acli/internal/prompt"

	"github.com/manifoldco/promptui"
)

var Pages = map[string]string{
	"terminal":  helprefs.Terminal,
	"rectangle": helprefs.Rectangle,
}

// GetReferencePage returns a string of the desired help text
func GetReferencePage(desiredPage string) (string, error) {
	if desiredPage != "" {
		return getPage(desiredPage)
	}

	pageName, err := promptPage()
	if err != nil {
		return "", fmt.Errorf("failed to prompt help page, %w", err)
	}
	return getPage(pageName)
}

// promptPage prompts the user to select a help page
func promptPage() (string, error) {
	var options []string
	for key := range Pages {
		options = append(options, key)
	}

	promptPage := promptui.Select{
		Label:        "Select Context",
		Items:        options,
		Stdout:       &prompt.BellSkipper{},
		Size:         15,
		HideHelp:     true,
		HideSelected: true,
	}

	_, pageName, err := promptPage.Run()
	if err != nil {
		return "", fmt.Errorf("failed to prompt the user to select a help page, %w", err)
	}

	return pageName, nil
}

// getPage grabs the help page from the internal/help/helprefs directory
func getPage(page string) (string, error) {
	helpPage := template.Must(template.New(page).Parse(Pages[page]))
	buf := &bytes.Buffer{}
	err := helpPage.Execute(buf, helprefs.TemplateData)
	if err != nil {
		return "", fmt.Errorf("failed to get help page, %w", err)
	}

	return buf.String(), nil
}
