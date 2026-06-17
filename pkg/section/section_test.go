package section

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func captureStdout(t *testing.T) (*bytes.Buffer, func()) {
	t.Helper()
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var buf bytes.Buffer
	done := make(chan struct{})
	go func() {
		io.Copy(&buf, r)
		close(done)
	}()

	return &buf, func() {
		w.Close()
		<-done
		os.Stdout = old
	}
}

func writeOverviewYML(t *testing.T, dir, content string) {
	t.Helper()
	path := dir + "/overview.yml"
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write overview.yml: %v", err)
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		rootDir string
		wantErr bool
	}{
		{
			name:    "valid root dir",
			rootDir: t.TempDir(),
			wantErr: false,
		},
		{
			name:    "empty root dir",
			rootDir: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := New(tt.rootDir)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("New() unexpected error: %v", err)
			}
			if s == nil {
				t.Fatal("expected non-nil Section")
			}
			if s.rootDir != tt.rootDir {
				t.Errorf("rootDir = %q, want %q", s.rootDir, tt.rootDir)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	dir := t.TempDir()
	writeOverviewYML(t, dir, `description:
  name: Test Chapter
  sections:
    - title: My Section
      text: My section content
`)

	s, err := New(dir)
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}

	buf, cleanup := captureStdout(t)
	err = s.Format("My Section")
	cleanup()

	if err != nil {
		t.Fatalf("Format() error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "My Section") {
		t.Error("output missing section title")
	}
	if !strings.Contains(output, "My section content") {
		t.Error("output missing section content")
	}
}

func TestFormatToString(t *testing.T) {
	dir := t.TempDir()
	writeOverviewYML(t, dir, `description:
  name: Test Chapter
  sections:
    - title: My Section
      text: My section content
`)

	s, err := New(dir)
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}

	got, err := s.FormatToString("My Section")
	if err != nil {
		t.Fatalf("FormatToString() error: %v", err)
	}

	if !strings.Contains(got, "My Section") {
		t.Error("output missing section title")
	}
	if !strings.Contains(got, "My section content") {
		t.Error("output missing section content")
	}
}

func TestFormatToStringEmptyTitle(t *testing.T) {
	dir := t.TempDir()
	s, err := New(dir)
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}

	_, err = s.FormatToString("")
	if err == nil {
		t.Fatal("expected error for empty title, got nil")
	}
}

func TestFormatToStringNotFound(t *testing.T) {
	dir := t.TempDir()
	writeOverviewYML(t, dir, `description:
  name: Test Chapter
  sections:
    - title: Existing
      text: Some content
`)

	s, err := New(dir)
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}

	_, err = s.FormatToString("Non-existent Section")
	if err == nil {
		t.Fatal("expected error for non-existent section, got nil")
	}
}

func TestValidateTitle(t *testing.T) {
	s := &Section{rootDir: t.TempDir()}

	if err := s.validateTitle("Valid Title"); err != nil {
		t.Errorf("expected nil for valid title, got: %v", err)
	}

	if err := s.validateTitle(""); err == nil {
		t.Error("expected error for empty title, got nil")
	}
}
