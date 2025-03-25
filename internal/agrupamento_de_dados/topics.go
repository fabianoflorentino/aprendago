/*
Package agrupamento_de_dados provides functionalities for demonstrating and executing
various data grouping topics in Go. It includes functions to print headers, execute
specific sections related to arrays, slices, and maps, and generate menu options and
help information for these topics. The package leverages the format package to format
and display the sections and help information.
*/
package agrupamento_de_dados

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"github.com/fabianoflorentino/aprendago/pkg/section"
	"github.com/fabianoflorentino/aprendago/pkg/topic"
)

// rootDir represents the relative path to the directory where data grouping topics are stored.
// This constant is used to reference the internal directory structure within the project.
const (
	rootDir = "internal/agrupamento_de_dados"
)

// Array represents a fixed-size sequence of elements of the same type.
// Arrays in Go have a fixed length, and their size is part of their type.
// They are useful when you know the exact number of elements you need to store.
const (
	Array                        string = "Array"
	SliceLiteralComposta         string = "Slice: literal composta"
	SliceForRange                string = "Slice: for range"
	SliceFatiandoOuDeletando     string = "Slice: fatiando ou deletando de uma fatia"
	SliceAnexando                string = "Slice: anexando a uma slice"
	SliceMake                    string = "Slice: make"
	SliceMultiDimensional        string = "Slice: multi dimensional"
	SliceSurpresaArraySubjacente string = "Slice: a surpresa do array subjacente"
	MapsIntroducao               string = "Maps: introdução"
	MapsRangeEDeletando          string = "Maps: range e deletando"
)

// AgrupamentoDeDados prints a header and executes a series of sections related to data grouping in Go.
// Each section is executed by calling the executeSection function with a specific topic name.
// The topics covered include arrays, slices (with various operations), and maps.
func AgrupamentoDeDados() {
	fmt.Printf("\n\n08 - Agrupamento de Dados\n")

	// listOfTopics is a slice of strings initialized with a length of 0 and a capacity of 10.
	// It is used to store a list of topics, allowing dynamic growth up to the specified capacity
	// without reallocating memory.
	listOfTopics := make([]string, 0, 10)
	listOfTopics = append(listOfTopics, // the first element is the list that you want to append new elements to it.
		Array,
		SliceLiteralComposta,
		SliceForRange,
		SliceFatiandoOuDeletando,
		SliceAnexando,
		SliceMake,
		SliceMultiDimensional,
		SliceSurpresaArraySubjacente,
		MapsIntroducao,
		MapsRangeEDeletando,
	)

	content := topic.New()
	content.TopicsContents(rootDir, listOfTopics)
}

// MenuAgrupamentoDeDados returns a slice of format.MenuOptions, each representing a menu option
// for different data grouping topics in Go. Each menu option includes a command-line option
// string and an associated function to execute a specific section related to arrays, slices,
// and maps. The function does not take any parameters and returns a slice of format.MenuOptions.
func MenuAgrupamentoDeDados([]string) []format.MenuOptions {
	section, err := section.New(rootDir)
	if err != nil {
		logger.New("Error creating section: ", err.Error()).RegisterLog()
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

// HelpMeAgrupamentoDeDados prints a list of topics related to data grouping in Go.
// It creates a slice of HelpMe structs, each containing a flag and a description
// of a specific topic. The topics include arrays, slices (with various operations),
// and maps. The function then prints the chapter title and uses the PrintHelpMe
// function from the format package to display the list of topics.
func HelpMeAgrupamentoDeDados() {
	hlp := []format.HelpMe{
		{Flag: "--array", Description: "Apresenta o tópico Array."},
		{Flag: "--slice-literal-composta", Description: "Apresenta o tópico Slice Literal Composta."},
		{Flag: "--slice-for-range", Description: "Apresenta o tópico Slice: for range."},
		{Flag: "--slice-fatiando-ou-deletando-de-uma-fatia", Description: "Apresenta o tópico Slice: fatiando ou deletando de uma fatia."},
		{Flag: "--slice-fatiando-ou-deletando-de-uma-fatia --resolucao", Description: "Apresenta a resolução do tópico Slice: fatiando ou deletando de uma fatia."},
		{Flag: "--slice-anexando-a-uma-slice", Description: "Apresenta o tópico Slice: anexando a uma slice."},
		{Flag: "--slice-make", Description: "Apresenta o tópico Slice: Make."},
		{Flag: "--slice-multi-dimensional", Description: "Apresenta o tópico Slice: Multi Dimensional."},
		{Flag: "--slice-a-surpresa-do-array-subjacente", Description: "Apresenta o tópico Slice: a surpresa do array subjacente."},
		{Flag: "--maps-introducao", Description: "Apresenta o tópico Maps: introdução."},
		{Flag: "--maps-range-e-deletando", Description: "Apresenta o tópico Maps: Range e Deletando."},
	}

	fmt.Println("\nCapítulo 8: Agrupamento de Dados")
	format.PrintHelpMe(hlp)
}
