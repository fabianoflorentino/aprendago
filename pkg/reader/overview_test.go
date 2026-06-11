package reader

import (
	"os"
	"path/filepath"
	"testing"
)

func writeOverviewYML(t *testing.T, dir, content string) {
	t.Helper()
	path := filepath.Join(dir, "overview.yml")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write overview.yml: %v", err)
	}
}

func writeOverviewYAML(t *testing.T, dir, content string) {
	t.Helper()
	path := filepath.Join(dir, "overview.yaml")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write overview.yaml: %v", err)
	}
}

func TestReadOverview(t *testing.T) {
	dir := t.TempDir()
	yml := `description:
  name: Test Chapter
  sections:
    - title: Section One
      text: Content one
    - title: Section Two
      text: Content two
`
	writeOverviewYML(t, dir, yml)

	docs, err := ReadOverview(dir)
	if err != nil {
		t.Fatalf("ReadOverview() error: %v", err)
	}

	if len(docs) != 1 {
		t.Fatalf("expected 1 document, got %d", len(docs))
	}

	if docs[0].Description.Name != "Test Chapter" {
		t.Errorf("expected name %q, got %q", "Test Chapter", docs[0].Description.Name)
	}

	if len(docs[0].Description.Sections) != 2 {
		t.Fatalf("expected 2 sections, got %d", len(docs[0].Description.Sections))
	}

	if docs[0].Description.Sections[0].Title != "Section One" {
		t.Errorf("expected section title %q, got %q", "Section One", docs[0].Description.Sections[0].Title)
	}
	if docs[0].Description.Sections[0].Text != "Content one" {
		t.Errorf("expected section text %q, got %q", "Content one", docs[0].Description.Sections[0].Text)
	}
	if docs[0].Description.Sections[1].Title != "Section Two" {
		t.Errorf("expected section title %q, got %q", "Section Two", docs[0].Description.Sections[1].Title)
	}
}

func TestReadOverviewYAMLExtension(t *testing.T) {
	dir := t.TempDir()
	yml := `description:
  name: YAML Chapter
  sections:
    - title: Only Section
      text: Some text
`
	writeOverviewYAML(t, dir, yml)

	docs, err := ReadOverview(dir)
	if err != nil {
		t.Fatalf("ReadOverview() error: %v", err)
	}

	if len(docs) != 1 {
		t.Fatalf("expected 1 document, got %d", len(docs))
	}

	if docs[0].Description.Name != "YAML Chapter" {
		t.Errorf("expected name %q, got %q", "YAML Chapter", docs[0].Description.Name)
	}
}

func TestReadOverviewPrefersYMLOverYAML(t *testing.T) {
	dir := t.TempDir()
	writeOverviewYML(t, dir, `description:
  name: From YML
  sections: []
`)
	writeOverviewYAML(t, dir, `description:
  name: From YAML
  sections: []
`)

	docs, err := ReadOverview(dir)
	if err != nil {
		t.Fatalf("ReadOverview() error: %v", err)
	}

	if docs[0].Description.Name != "From YML" {
		t.Errorf("expected to prefer overview.yml, got %q", docs[0].Description.Name)
	}
}

func TestReadOverviewInvalidYAML(t *testing.T) {
	dir := t.TempDir()
	writeOverviewYML(t, dir, `invalid yaml: [`)
	_, err := ReadOverview(dir)
	if err == nil {
		t.Fatal("expected error for invalid YAML, got nil")
	}
}

func TestReadSection(t *testing.T) {
	dir := t.TempDir()
	writeOverviewYML(t, dir, `description:
  name: Test Chapter
  sections:
    - title: Target Section
      text: Target content
    - title: Other Section
      text: Other content
`)

	sec, err := ReadSection(dir, "Target Section")
	if err != nil {
		t.Fatalf("ReadSection() error: %v", err)
	}

	if sec.Title != "Target Section" {
		t.Errorf("expected title %q, got %q", "Target Section", sec.Title)
	}
	if sec.Text != "Target content" {
		t.Errorf("expected text %q, got %q", "Target content", sec.Text)
	}
}

func TestReadSectionNotFound(t *testing.T) {
	dir := t.TempDir()
	writeOverviewYML(t, dir, `description:
  name: Test Chapter
  sections:
    - title: Existing Section
      text: Some text
`)

	_, err := ReadSection(dir, "Non-existent Section")
	if err == nil {
		t.Fatal("expected error for non-existent section, got nil")
	}
	if !os.IsNotExist(err) {
		t.Errorf("expected os.ErrNotExist, got %v", err)
	}
}

