package common

import (
	"fmt"
	"os"
	"os/exec"
)

// PrintStringWithMore will write text to a temp file and read it with the `less -mX FILENAME` cmd.
func PrintStringWithMore(text string) error {
	tempFile, err := os.CreateTemp("", "print_with_more_temp_file.txt")
	if err != nil {
		return fmt.Errorf("failed to create temp file, %w", err)
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString(text)
	if err != nil {
		return fmt.Errorf("failed to write text to temp file, %w", err)
	}

	err = tempFile.Close()
	if err != nil {
		return fmt.Errorf("failed to close temp file, %w", err)
	}

	moreCmd := exec.Command(
		"less",
		"-mX",
		tempFile.Name(),
	)
	moreCmd.Stdout = os.Stdout
	err = moreCmd.Run()
	if err != nil {
		return fmt.Errorf("failed to print out text with the more command, %w", err)
	}
	return nil
}
