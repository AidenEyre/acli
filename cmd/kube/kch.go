package cmd

import (
	"fmt"
	"os/exec"

	"github.com/aideneyre/acli/internal/config"
	"github.com/aideneyre/acli/internal/kube"

	"github.com/spf13/cobra"
)

var (
	resourceType  string
	namespace     string
	KubeChangeCmd = &cobra.Command{
		Use:   "kch [flags] [context]",
		Short: "Change kube context.",
		Long:  "Change kube context. If kch has been configured, may open k9s by default to a specific resource type and namespace.",
		Example: `
# Set context with flag
acli kch CONTEXT_NAME

# Set the context with the help of an interactive prompt
acli kch

# Set the context with the help of an interactive prompt and open K9s to deployments in the
# default namespace
acli kch -r deploy

# Set the context with the help of an interactive prompt and open K9s to pods in the kube-system
# namespace
acli kch -r po -n kube-system
		`,
		Version:      "2.2.1",
		RunE:         runE,
		SilenceUsage: true,
	}
)

func init() {
	KubeChangeCmd.Flags().StringVarP(&resourceType, "resource", "r", "", "Specify the resource type")
	KubeChangeCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Specify the namespace")
}

// runE calls sets current kube context with a prompt or a specified context.
func runE(cmd *cobra.Command, args []string) error {
	var err error
	err = config.Initialize()
	if err != nil {
		return fmt.Errorf("unable to initialize config: %w", err)
	}

	if len(args) > cmd.Flags().NFlag() {
		desiredContext := args[0]
		err = kube.SetKubeContext(desiredContext)
	} else {
		err = kube.SetKubeContextWithPrompt()
	}
	if err != nil {
		return fmt.Errorf("unable to set context: %w", err)
	}

	// Stop here if not opening k9s
	if resourceType == "" && namespace == "" && !config.GetBool("kch.alwaysopenk9s") {
		return nil
	}

	_, err = exec.LookPath("k9s")
	if err != nil {
		fmt.Println("k9s is not installed or not found in $PATH. You must install k9s to open it. For installation instructions, visit https://k9scli.io/topics/install/")
		return nil
	}

	if namespace == "" {
		namespace = config.GetString("kch.defaults.namespace")
	}
	if resourceType == "" {
		resourceType = config.GetString("kch.defaults.resource")
	}

	err = kube.RunK9s(resourceType, namespace)
	if err != nil {
		return fmt.Errorf("failed to run K9s: %+w", err)
	}
	return nil
}
