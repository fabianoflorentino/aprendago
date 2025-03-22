// Package section provides functionality to handle and format sections of documents.
// It includes methods to read sections from a directory and format them for overview purposes.
package section

import (
	"fmt"
	"os"

	"github.com/fabianoflorentino/aprendago/pkg/format"
	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"github.com/fabianoflorentino/aprendago/pkg/reader"
)

// Section represents a part or division of a larger structure or document.
// It is currently an empty struct, but can be extended with fields and methods
// to hold and manipulate section-specific data.
type Section struct {
	rootDir string
}

// New creates and returns a new instance of Section.
func New(rootDir string) (*Section, error) {
	if rootDir == "" {
		return nil, fmt.Errorf("diretório raiz não pode ser vazio")
	}
	return &Section{rootDir: rootDir}, nil
}

// Format reads a section from the specified directory and title, and formats it.
// If the section is not found, it logs an appropriate message.
// If there is an error reading the section or formatting the document, it logs the error.
//
// Parameters:
//
//	dir: The directory where the section is located.
//	title: The title of the section to be read.
func (s *Section) Format(title string) error {
	if err := s.validateTitle(title); err != nil {
		logger.Log("Erro ao validar o título da seção: %v", err)
		return err
	}

	sec, err := s.readSection(title)
	if err != nil {
		logger.Log("Erro ao ler a seção: %v", err)
		return err
	}

	return s.formatDocument(sec)
}

// validateTitle checks if the provided title is not empty.
// It returns an error if the title is empty, otherwise it returns nil.
//
// Parameters:
//   - title: The title string to be validated.
//
// Returns:
//   - error: An error if the title is empty, otherwise nil.
func (s *Section) validateTitle(title string) error {
	if title == "" {
		return fmt.Errorf("título da seção não pode ser vazio")
	}
	return nil
}

// readSection reads a section with the given title from the root directory of the Section.
// It returns a pointer to the reader.Section and an error if any occurs during the reading process.
//
// Parameters:
//   - title: The title of the section to be read.
//
// Returns:
//   - *reader.Section: A pointer to the read section.
//   - error: An error if the section is not found or if there is an error during the reading process.
//
// If the section is not found, it logs the error and returns a formatted error message indicating
// that the section with the given title was not found in the specified root directory.
// If there is any other error during the reading process, it logs the error and returns a generic
// error message indicating that there was an error reading the section.
func (s *Section) readSection(title string) (*reader.Section, error) {
	sec, err := reader.ReadSection(s.rootDir, title)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Log("Seção '%s' não encontrada no diretório '%s'", title, s.rootDir)
			return nil, fmt.Errorf("seção '%s' não encontrada no diretório '%s'", title, s.rootDir)
		}

		logger.Log("Erro ao ler a seção: %v", err)
		return nil, fmt.Errorf("erro ao ler a seção")
	}

	if sec == nil {
		logger.Log("Seção '%s' não encontrada no diretório '%s'", title, s.rootDir)
		return nil, fmt.Errorf("seção não encontrada")
	}

	return sec, nil
}

// formatDocument formats a given section into a document and processes it using the format.FormatOverview function.
// If an error occurs during formatting, it logs the error and returns a wrapped error.
//
// Parameters:
//   - sec: A pointer to a reader.Section that needs to be formatted.
//
// Returns:
//   - error: An error if the formatting fails, otherwise nil.
func (s *Section) formatDocument(sec *reader.Section) error {
	document := reader.Document{
		Description: reader.Description{
			Sections: []reader.Section{*sec},
		},
	}

	err := format.FormatOverview([]reader.Document{document})
	if err != nil {
		logger.Log("Erro ao formatar o documento: %v", err)
		return fmt.Errorf("erro ao formatar o documento: %w", err)
	}

	return nil
}
