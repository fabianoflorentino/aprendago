// Package reader provides functionality to read and parse overview files in YAML format.
// The overview files contain descriptions and sections, each with a title and text.
// This package includes functions to read the entire overview document or specific sections by title.
// The overview file can be named either overview.yml or overview.yaml.
package reader

import (
	"os"
	"path/filepath"

	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"gopkg.in/yaml.v2"
)

// DOCTOREAD_YML represents the filename for the overview document in YAML format with the .yml extension.
const (
	DOCTOREAD_YML  = "overview.yml"
	DOCTOREAD_YAML = "overview.yaml"
)

// Section represents a section of the overview
//   - Title: the title of the section
//   - Text: the text of the section
type Section struct {
	Title string `yaml:"title"`
	Text  string `yaml:"text"`
}

// Description represents the description of the overview
//   - Name: the name of the overview
//   - Sections: the sections of the overview
type Description struct {
	Name     string    `yaml:"name"`
	Sections []Section `yaml:"sections"`
}

// Document represents the document of the overview
//   - Description: the description of the overview
type Document struct {
	Description Description `yaml:"description"`
}

// ReadSection reads a section from the root directory
//   - rootDir: the root directory of the overview
//   - sectionTitle: the title of the section
//
// Returns the section of the overview
// Returns an error if the section is not found
func ReadSection(rootDir, sectionTitle string) (*Section, error) {
	documents, err := readOverview(rootDir)
	if err != nil {
		return nil, err
	}

	for _, document := range documents {
		for _, section := range document.Description.Sections {
			if section.Title == sectionTitle {
				return &section, nil
			}
		}
	}

	return nil, os.ErrNotExist
}

// ReadOverview reads the overview from the root directory
//   - rootDir: the root directory of the overview
//
// Returns the document of the overview
// Returns an error if the overview file is not found or if the file is not valid
// The overview file can be either overview.yml or overview.yaml
//   - overview.yml: the YAML file of the overview
//   - overview.yaml: the YAML file of the overview
//   - docPath: the path of the overview file
func readOverview(rootDir string) ([]Document, error) {
	var documentSpecification []Document
	var specification Document
	var docPath string

	docPathYML := filepath.Join(rootDir, DOCTOREAD_YML)
	docPathYAML := filepath.Join(rootDir, DOCTOREAD_YAML)

	if _, err := os.Stat(docPathYML); err == nil {
		docPath = docPathYML
	} else if _, err := os.Stat(docPathYAML); err == nil {
		docPath = docPathYAML
	} else {
		logger.Log("File %s or %s not found", docPathYML, docPathYAML)
		return nil, os.ErrNotExist
	}

	data, err := os.ReadFile(docPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &specification)
	if err != nil {
		return nil, err
	}

	documentSpecification = append(documentSpecification, specification)

	return documentSpecification, nil
}
