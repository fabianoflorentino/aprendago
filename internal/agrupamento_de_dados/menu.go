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
	}

	return []format.MenuOptions{
		{Options: "--array", ExecFunc: func() { sectionsAgrupamentoDeDados(array, section) }},
		{Options: "--slice-literal-composta", ExecFunc: func() { sectionsAgrupamentoDeDados(sliceLiteralComposta, section) }},
		{Options: "--slice-for-range", ExecFunc: func() { sectionsAgrupamentoDeDados(sliceForRange, section) }},
		{Options: "--slice-fatiando-ou-deletando-de-uma-fatia", ExecFunc: func() { sectionsAgrupamentoDeDados(sliceFatiandoOuDeletando, section) }},
		{Options: "--slice-anexando-a-uma-slice", ExecFunc: func() { sectionsAgrupamentoDeDados(sliceAnexando, section) }},
		{Options: "--slice-make", ExecFunc: func() { sectionsAgrupamentoDeDados(sliceMake, section) }},
		{Options: "--slice-multi-dimensional", ExecFunc: func() { sectionsAgrupamentoDeDados(sliceMultiDimensional, section) }},
		{Options: "--slice-a-surpresa-do-array-subjacente", ExecFunc: func() { sectionsAgrupamentoDeDados(sliceSurpresaArraySubjacente, section) }},
		{Options: "--maps-introducao", ExecFunc: func() { sectionsAgrupamentoDeDados(mapsIntroducao, section) }},
		{Options: "--maps-range-e-deletando", ExecFunc: func() { sectionsAgrupamentoDeDados(mapsRangeEDeletando, section) }},
	}
}

// sectionsAgrupamentoDeDados formats and processes sections of data under a given title.
// It utilizes the SectionProvider interface to handle the formatting logic.
//
// Parameters:
//   - title: A string representing the title for the sections.
//   - sections: An implementation of the section.SectionProvider interface
//     that provides the functionality to format the sections.
func sectionsAgrupamentoDeDados(title string, sections section.SectionProvider) {
	sections.Format(title)
}
