package connectssh

import (
	"fmt"
	"os"
	"os/exec"
	"sort"

	connectConfig "github.com/aideneyre/acli/internal/config/configurator/connect"
	"github.com/aideneyre/acli/internal/prompt"
	"github.com/manifoldco/promptui"
)

func Connect(globalConfig connectConfig.Config, serverAlias string) error {
	var selectedServerIndex int
	for index, alias := range globalConfig.Connect.Aliases {
		if alias.Name == serverAlias {
			selectedServerIndex = index
			break
		}
	}
	server := globalConfig.Connect.Aliases[selectedServerIndex]

	privateKeyPath := fmt.Sprintf("%s/%s", globalConfig.Connect.SSHDir, server.PEMFile)

	cmd := exec.Command("ssh",
		"-i",
		privateKeyPath,
		fmt.Sprintf("%s@%s", server.User, server.IP),
		"-p",
		"22",
	)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("unable to connect to server: %w", err)
	}

	return nil
}

func ConnectInteractive(globalConfig connectConfig.Config) error {
	var serverNames []string
	for _, alias := range globalConfig.Connect.Aliases {
		serverNames = append(serverNames, alias.Name)
	}

	selectedServer, err := promptServer(serverNames)
	if err == promptui.ErrInterrupt {
		fmt.Print("\033[u\033[J") // Clear the terminal
		fmt.Println("CLI exited early!")
		os.Exit(0)
	}
	if err != nil {
		return fmt.Errorf("unable to get user context selection: %w", err)
	}

	return Connect(globalConfig, selectedServer)
}

func promptServer(servers []string) (string, error) {
	sort.Strings(servers)
	promptServer := promptui.Select{
		Label:        "Select server",
		Items:        servers,
		Stdout:       &prompt.BellSkipper{},
		Size:         15,
		HideHelp:     true,
		HideSelected: true,
	}

	_, server, err := promptServer.Run()
	if err != nil {
		return "", err
	}

	return server, nil
}
