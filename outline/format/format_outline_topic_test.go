package format

import (
	"io"
	"os"
	"testing"
)

func TestFormatOutlineTopic(t *testing.T) {
	topic := OutlineContent{
		Title:   "Sample Title",
		Content: "Sample Content",
	}

	expectedOutput := "Sample Title \n Sample Content\n"
	if got := outlineTopiccaptureOutput(FormatOutlineTopic, topic); got != expectedOutput {
		t.Errorf("Expected output: %q, but got: %q", expectedOutput, got)
	}
}

// Helper function to capture the output of a function
func outlineTopiccaptureOutput(fn func(OutlineContent), topic OutlineContent) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fn(topic)
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = old
	return string(out)
}
