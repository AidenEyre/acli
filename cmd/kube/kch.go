package cmd

import (
	"fmt"

	"github.com/aideneyre/acli/internal/kube"

	"github.com/spf13/cobra"
)

var (
	resourceType  string
	namespace     string
	KubeChangeCmd = &cobra.Command{
		Use:   "kch [flags] [context]",
		Short: "Change kube context.",
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
		Version:      "2.1.0",
		RunE:         runE,
		SilenceUsage: true,
	}
)

func init() {
	KubeChangeCmd.Flags().StringVarP(&resourceType, "resource", "r", "", "Specify the resource type")
	KubeChangeCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "Specify the namespace")
	KubeChangeCmd.Flags().Lookup("resource").NoOptDefVal = "pod"
	KubeChangeCmd.Flags().Lookup("namespace").NoOptDefVal = "default"
}

// runE calls sets current kube context with a prompt or a specified context.
func runE(cmd *cobra.Command, args []string) error {
	var err error

	if len(args) > 0 {
		desiredContext := args[0]
		err = kube.SetKubeContext(desiredContext)
	} else {
		err = kube.SetKubeContextWithPrompt()
	}
	if err != nil {
		return fmt.Errorf("unable to set context: %w", err)
	}

	if resourceType != "" || namespace != "" {
		err = kube.RunK9s(resourceType, namespace)
	}
	if err != nil {
		return fmt.Errorf("failed to run K9s: %+w", err)
	}

	return nil
}
