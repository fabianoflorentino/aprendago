// Package visao_geral_do_curso provides an overview of the course "Aprenda Go".
// It includes functions to display different sections of the course,
// generate a menu with options to navigate through the sections,
// and provide help descriptions for each section.
package visao_geral_do_curso

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// rootDir represents the root directory path for the "visao_geral_do_curso" module within the project.
const (
	rootDir = "internal/visao_geral_do_curso"
)

// VisaoGeralDoCurso prints an overview of the course topics to the console.
// It sequentially executes sections that welcome the user, explain why Go is beneficial,
// outline success strategies, provide resources, and describe how the course works.
func VisaoGeralDoCurso() {
	fmt.Printf("\n01 - Visão Geral do Curso\n")

	executeSection("Bem-vindo!")
	executeSection("Por que Go?")
	executeSection("Sucesso")
	executeSection("Recursos")
	executeSection("Como esse curso funciona")
}

// MenuVisaoGeralDoCurso returns a slice of MenuOptions for the course overview menu.
// Each MenuOption contains an option string and an associated function to execute when the option is selected.
// The options include:
// --bem-vindo: Executes the "Bem-vindo!" section.
// --porque-go: Executes the "Por que Go?" section.
// --sucesso: Executes the "Sucesso" section.
// --recursos: Executes the "Recursos" section.
// --como-esse-curso-funciona: Executes the "Como esse curso funciona" section.
func MenuVisaoGeralDoCurso([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--bem-vindo", ExecFunc: func() { executeSection("Bem-vindo!") }},
		{Options: "--porque-go", ExecFunc: func() { executeSection("Por que Go?") }},
		{Options: "--sucesso", ExecFunc: func() { executeSection("Sucesso") }},
		{Options: "--recursos", ExecFunc: func() { executeSection("Recursos") }},
		{Options: "--como-esse-curso-funciona", ExecFunc: func() { executeSection("Como esse curso funciona") }},
	}
}

// HelpMeVisaoGeralDoCurso provides a list of help topics related to the overview of the "Aprenda Go" course.
// It includes flags and descriptions for various aspects such as welcome message, benefits of learning Go,
// tips for success, available resources, and course structure.
func HelpMeVisaoGeralDoCurso() {
	hlp := []format.HelpMe{
		{Flag: "--bem-vindo", Description: "Exibe a mensagem de boas-vindas ao curso Aprenda Go.", Width: 0},
		{Flag: "--porque-go", Description: "Descreve os benefícios e razões para aprender a linguagem Go.", Width: 0},
		{Flag: "--sucesso", Description: "Apresenta dicas e estratégias para ter sucesso no curso.", Width: 0},
		{Flag: "--recursos", Description: "Lista recursos e materiais de apoio para o curso.", Width: 0},
		{Flag: "--como-esse-curso-funciona", Description: "Explica a estrutura e metodologia do curso.", Width: 0},
	}

	fmt.Println("\nCapítulo 1: Visão Geral do Curso")
	format.PrintHelpMe(hlp)
}

// executeSection formats and processes a given section of the course.
// It takes the section name as a string parameter and uses the FormatSection
// function from the format package to apply the formatting to the section
// located in the root directory.
func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
