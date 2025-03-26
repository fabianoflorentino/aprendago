package agrupamento_de_dados

import (
	"github.com/fabianoflorentino/aprendago/pkg/format"
	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"github.com/fabianoflorentino/aprendago/pkg/section"
)

// menu generates a list of menu options based on the provided menu slice and root directory.
// It initializes a section using the given root directory and returns a slice of MenuOptions,
// each containing an option string and an associated execution function.
//
// Parameters:
//   - menu: A slice of strings representing the menu options.
//   - rootDir: A string representing the root directory for initializing the section.
//
// Returns:
//   - A slice of format.MenuOptions, where each option has an associated execution function.
//     If an error occurs during section initialization, it logs the error and returns nil.
func Menu(menu []string) []format.MenuOptions {
	section, err := section.New(rootDir)
	if err != nil {
		logger.New("Error creating section: ", err.Error()).Register()
		return nil
	}

	return []format.MenuOptions{
		{Options: "--array", ExecFunc: func() { section.Format("Array") }},
		{Options: "--slice-literal-composta", ExecFunc: func() { section.Format("Slice: literal composta") }},
		{Options: "--slice-for-range", ExecFunc: func() { section.Format("Slice: for range") }},
		{Options: "--slice-fatiando-ou-deletando-de-uma-fatia", ExecFunc: func() { section.Format("Slice: fatiando ou deletando de uma fatia") }},
		{Options: "--slice-anexando-a-uma-slice", ExecFunc: func() { section.Format("Slice: anexando a uma slice") }},
		{Options: "--slice-make", ExecFunc: func() { section.Format("Slice: make") }},
		{Options: "--slice-multi-dimensional", ExecFunc: func() { section.Format("Slice: multi dimensional") }},
		{Options: "--slice-a-surpresa-do-array-subjacente", ExecFunc: func() { section.Format("Slice: a surpresa do array subjacente") }},
		{Options: "--maps-introducao", ExecFunc: func() { section.Format("Maps: introdução") }},
		{Options: "--maps-range-e-deletando", ExecFunc: func() { section.Format("Maps: range e deletando") }},
	}
}
