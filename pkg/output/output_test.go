package output

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	o := New()
	if o == nil {
		t.Fatal("expected non-nil *Output from New()")
	}
	if o.stdout != os.Stdout {
		t.Error("stdout field should be os.Stdout")
	}
	if o.stderr != os.Stderr {
		t.Error("stderr field should be os.Stderr")
	}
	if o.read == nil {
		t.Error("read end of pipe should not be nil")
	}
	if o.write == nil {
		t.Error("write end of pipe should not be nil")
	}
}

func TestCapture(t *testing.T) {
	o := New()
	msg := "hello from capture test"

	got, err := o.Capture(func() {
		fmt.Print(msg)
	})
	if err != nil {
		t.Fatalf("Capture() returned error: %v", err)
	}

	if strings.TrimSpace(got) != msg {
		t.Errorf("Capture() = %q, want %q", strings.TrimSpace(got), msg)
	}

	// Verify stdout is restored
	if os.Stdout != o.stdout {
		t.Error("stdout not restored after Capture")
	}
}

func TestCaptureEmptyOutput(t *testing.T) {
	o := New()
	got, err := o.Capture(func() {
		// do nothing
	})
	if err != nil {
		t.Fatalf("Capture() returned error: %v", err)
	}
	if got != "" {
		t.Errorf("Capture() = %q, want empty string", got)
	}
}

func TestCaptureMultiple(t *testing.T) {
	o1 := New()
	got, err := o1.Capture(func() {
		fmt.Print("first")
	})
	if err != nil {
		t.Fatalf("first Capture() error: %v", err)
	}
	if strings.TrimSpace(got) != "first" {
		t.Errorf("first Capture() = %q, want %q", got, "first")
	}

	o2 := New()
	got2, err := o2.Capture(func() {
		fmt.Print("second")
	})
	if err != nil {
		t.Fatalf("second Capture() error: %v", err)
	}
	if strings.TrimSpace(got2) != "second" {
		t.Errorf("second Capture() = %q, want %q", got2, "second")
	}
}
