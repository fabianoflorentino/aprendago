package format

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestFormatMenuOptions(t *testing.T) {
	args := []string{"--help"}
	options := []MenuOptions{
		{
			Options: "--help",
			ExecFunc: func() {
				fmt.Println("Help option selected")
			},
		},
		{
			Options: "--version",
			ExecFunc: func() {
				fmt.Println("Version option selected")
			},
		},
	}
	FormatMenuOptions(args, options)
	// Assertion to validate the expected output
	expectedOutput := "Help option selected\n"
	if got := captureOutput(FormatMenuOptions, args, options); got != expectedOutput {
		t.Errorf("Expected output: %q, but got: %q", expectedOutput, got)
	}
}

// Helper function to capture the output of a function
func captureOutput(fn func([]string, []MenuOptions), args []string, options []MenuOptions) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn(args, options)

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = old

	return string(out)
}
