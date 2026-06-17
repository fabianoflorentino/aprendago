package format

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/reader"
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

func TestFormatOverview(t *testing.T) {
	docs := []reader.Document{
		{
			Description: reader.Description{
				Name: "Test Chapter",
				Sections: []reader.Section{
					{Title: "Section 1", Text: "Content 1"},
					{Title: "Section 2", Text: "Content 2"},
				},
			},
		},
	}

	buf, cleanup := captureStdout(t)
	err := FormatOverview(docs)
	cleanup()

	if err != nil {
		t.Fatalf("FormatOverview() returned error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "Test Chapter") {
		t.Error("output missing chapter name")
	}
	if !strings.Contains(output, "Section 1") {
		t.Error("output missing section 1 title")
	}
	if !strings.Contains(output, "Content 1") {
		t.Error("output missing section 1 content")
	}
	if !strings.Contains(output, "Section 2") {
		t.Error("output missing section 2 title")
	}
	if !strings.Contains(output, "Content 2") {
		t.Error("output missing section 2 content")
	}
}

func TestFormatOverviewSingleSection(t *testing.T) {
	docs := []reader.Document{
		{
			Description: reader.Description{
				Name: "Single Section Chapter",
				Sections: []reader.Section{
					{Title: "Only Section", Text: "Only content"},
				},
			},
		},
	}

	buf, cleanup := captureStdout(t)
	err := FormatOverview(docs)
	cleanup()

	if err != nil {
		t.Fatalf("FormatOverview() returned error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "Single Section Chapter") {
		t.Error("output missing chapter name")
	}
	if !strings.Contains(output, "Only Section") {
		t.Error("output missing section title")
	}
	if !strings.Contains(output, "Only content") {
		t.Error("output missing section content")
	}
}

func TestFormatOverviewMultipleDocuments(t *testing.T) {
	docs := []reader.Document{
		{
			Description: reader.Description{
				Name: "Chapter One",
				Sections: []reader.Section{
					{Title: "Topic 1", Text: "Text 1"},
				},
			},
		},
		{
			Description: reader.Description{
				Name: "Chapter Two",
				Sections: []reader.Section{
					{Title: "Topic 2", Text: "Text 2"},
				},
			},
		},
	}

	buf, cleanup := captureStdout(t)
	err := FormatOverview(docs)
	cleanup()

	if err != nil {
		t.Fatalf("FormatOverview() returned error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "Chapter One") {
		t.Error("output missing first document name")
	}
	if !strings.Contains(output, "Chapter Two") {
		t.Error("output missing second document name")
	}
}

func TestFormatOverviewTo(t *testing.T) {
	docs := []reader.Document{
		{
			Description: reader.Description{
				Name: "Test Chapter",
				Sections: []reader.Section{
					{Title: "Section 1", Text: "Content 1"},
				},
			},
		},
	}

	var buf bytes.Buffer
	err := FormatOverviewTo(&buf, docs)

	if err != nil {
		t.Fatalf("FormatOverviewTo() returned error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "Test Chapter") {
		t.Error("output missing chapter name")
	}
	if !strings.Contains(output, "Section 1") {
		t.Error("output missing section title")
	}
	if !strings.Contains(output, "Content 1") {
		t.Error("output missing section content")
	}
}

func TestFormatOverviewToEmptyDocuments(t *testing.T) {
	var buf bytes.Buffer
	err := FormatOverviewTo(&buf, []reader.Document{})

	if err != nil {
		t.Fatalf("FormatOverviewTo() returned error: %v", err)
	}
}

func TestFormatOverviewToMultipleDocuments(t *testing.T) {
	docs := []reader.Document{
		{
			Description: reader.Description{
				Name: "Doc One",
				Sections: []reader.Section{
					{Title: "T1", Text: "Text 1"},
				},
			},
		},
		{
			Description: reader.Description{
				Name: "Doc Two",
				Sections: []reader.Section{
					{Title: "T2", Text: "Text 2"},
				},
			},
		},
	}

	var buf bytes.Buffer
	err := FormatOverviewTo(&buf, docs)
	if err != nil {
		t.Fatalf("FormatOverviewTo() returned error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "Doc One") {
		t.Error("output missing first document name")
	}
	if !strings.Contains(output, "Doc Two") {
		t.Error("output missing second document name")
	}
}

func TestFormatOverviewNoSections(t *testing.T) {
	docs := []reader.Document{
		{
			Description: reader.Description{
				Name:     "Empty Chapter",
				Sections: []reader.Section{},
			},
		},
	}

	buf, cleanup := captureStdout(t)
	err := FormatOverview(docs)
	cleanup()

	if err != nil {
		t.Fatalf("FormatOverview() returned error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "Empty Chapter") {
		t.Error("output missing chapter name")
	}
}
