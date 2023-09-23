package kube

import (
	"fmt"
	"os"
	"os/exec"
)

// RunK9s will open up k9s to the specified resource type and namespace.
func RunK9s(resourceType, namespace string) error {
	cmd := exec.Command("k9s",
		fmt.Sprintf("--namespace=%s", namespace),
		fmt.Sprintf("--command=%s", resourceType))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
