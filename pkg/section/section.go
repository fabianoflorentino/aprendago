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
	if title == "" {
		return fmt.Errorf("título da seção não pode ser vazio")
	}

	section, err := reader.ReadSection(s.rootDir, title)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Log("Seção '%s' não encontrada no diretório '%s'", title, s.rootDir)
		}

		logger.Log("Erro ao ler a seção: %v", err)
		return err
	}

	if section == nil {
		logger.Log("Seção '%s' não encontrada no diretório '%s'", title, s.rootDir)
		return fmt.Errorf("seção não encontrada")
	}

	document := reader.Document{
		Description: reader.Description{
			Sections: []reader.Section{*section},
		},
	}

	err = format.FormatOverview([]reader.Document{document})
	if err != nil {
		logger.Log("Erro ao formatar o documento: %v", err)
		return fmt.Errorf("erro ao formatar o documento")
	}

	return nil
}
