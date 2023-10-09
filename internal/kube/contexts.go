// Package kube offers utilities for working with Kubernetes.
package kube

import (
	"fmt"
	"log"
	"sort"

	"github.com/aideneyre/acli/internal/common"
	"github.com/aideneyre/acli/internal/prompt"

	"github.com/manifoldco/promptui"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

var (
	kubeConfig     *api.Config
	kubeConfigPath string
)

func init() {
	var err error
	kubeConfig, kubeConfigPath, err = loadConfig()
	if err != nil {
		log.Fatalf("failed to initialize K8s API config:: %v", err)
	}
}

// SetKubeContext takes in the context name and sets it as the current kube context.
func SetKubeContext(context string) error {
	kubeConfig.CurrentContext = context
	err := clientcmd.WriteToFile(*kubeConfig, kubeConfigPath)
	if err != nil {
		return err
	}
	fmt.Printf("context successfully set to %s\n", context)
	return nil
}

// SetKubeContextWithPrompt will grab all the available contexts, prompt the user to make a choice,
// and set the context.
func SetKubeContextWithPrompt() error {
	currentContext := kubeConfig.CurrentContext
	contexts := make([]string, 0, len(kubeConfig.Contexts))
	for ctx := range kubeConfig.Contexts {
		contexts = append(contexts, ctx)
	}

	if len(contexts) == 0 {
		return fmt.Errorf("failed to find configured contexts, nothing to prompt")
	}

	contexts = common.ColorSliceStringGreen(contexts, currentContext)
	selectedContext, err := promptContext(contexts)
	if err == promptui.ErrInterrupt {
		fmt.Print("\033[u\033[J") // Clear the terminal
		fmt.Println("CLI exited early!")
		return nil
	}
	if err != nil {
		return fmt.Errorf("unable to get user context selection: %w", err)
	}
	selectedContext = common.RemoveGreenStringFormatting(selectedContext)

	err = SetKubeContext(selectedContext)
	if err != nil {
		return fmt.Errorf("failed to write kube config: %w", err)
	}

	return nil
}

// loadConfig returns a Kubernetes client API config and the defined config
// path
func loadConfig() (*api.Config, string, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	overrides := &clientcmd.ConfigOverrides{}

	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, overrides)
	apiConfig, err := clientConfig.RawConfig()
	if err != nil {
		return nil, "", err
	}

	// Use the first file if multiple exist, as it's the highest precedence
	if len(loadingRules.Precedence) == 0 {
		return nil, "", fmt.Errorf("failed to find config file")
	}
	return &apiConfig, loadingRules.Precedence[0], nil
}

func promptContext(contexts []string) (string, error) {
	sort.Strings(contexts)
	promptCtx := promptui.Select{
		Label:        "Select Context",
		Items:        contexts,
		Stdout:       &prompt.BellSkipper{},
		Size:         15,
		HideHelp:     true,
		HideSelected: true,
	}

	_, context, err := promptCtx.Run()
	if err != nil {
		return "", err
	}

	return context, nil
}
