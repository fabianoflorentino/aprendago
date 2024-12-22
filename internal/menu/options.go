/*
Package menu provides a command-line interface for navigating the options and functionalities
of the Learn GO course. It groups and organizes different menus and their options, allowing the user
to access specific chapters, practical exercises, and course overviews based on the provided arguments.

Main functionalities:

- Displays the main menu with general and chapter-specific options.
- Allows access to different sections of the course, such as variables, types, control flow, and practical exercises.
- Processes the arguments provided by the user to determine and execute the appropriate option.
- Supports help and overview options to facilitate navigation.

Basic usage:

To use the menu package, you should call the Options function with a list of arguments.
These arguments determine which menu or option will be displayed and executed.

Example:

	args := []string{"--help"}
	menu.Options(args)

In this example, the Options function will display the help, based on the provided argument.
*/
package menu

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/internal"
	"github.com/fabianoflorentino/aprendago/internal/agrupamento_de_dados"
	"github.com/fabianoflorentino/aprendago/internal/aplicacoes"
	"github.com/fabianoflorentino/aprendago/internal/canais"
	"github.com/fabianoflorentino/aprendago/internal/concorrencia"
	"github.com/fabianoflorentino/aprendago/internal/documentacao"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_1"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_10"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_11"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_12"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_2"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_3"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_4"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_5"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_6"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_7"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_8"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_9"
	"github.com/fabianoflorentino/aprendago/internal/fluxo_de_controle"
	"github.com/fabianoflorentino/aprendago/internal/funcoes"
	"github.com/fabianoflorentino/aprendago/internal/fundamentos_da_programacao"
	"github.com/fabianoflorentino/aprendago/internal/ponteiros"
	"github.com/fabianoflorentino/aprendago/internal/seu_ambiente_de_desenvolvimento"
	"github.com/fabianoflorentino/aprendago/internal/structs"
	"github.com/fabianoflorentino/aprendago/internal/tratamento_de_erro"
	"github.com/fabianoflorentino/aprendago/internal/variaveis_valores_tipos"
	"github.com/fabianoflorentino/aprendago/internal/visao_geral_do_curso"
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// Options displays the main menu with general and chapter-specific options based on the provided arguments.
// It allows the user to navigate through the course content, access practical exercises, and view course overviews.
//   - args: A list of arguments provided by the user to determine the menu or option to be displayed.
//
// The Options function calls the following functions:
//   - buildOptions: Combines general and chapter-specific options based on the provided arguments.
//   - generalOptions: Returns the general options available in the main menu.
func Options(args []string) {
	fmt.Printf("Learn GO\n\n")

	buildOptions(
		args,
		generalOptions(args),
		MenuCapituloOptions(args),
		MenuCapituloOutline(args),
		visao_geral_do_curso.MenuVisaoGeralDoCurso(args),
		variaveis_valores_tipos.MenuVariaveisValoresTipos(args),
		exercicios_ninja_nivel_1.MenuExerciciosNinjaNivel1(args),
		fundamentos_da_programacao.MenuFundamentosDaProgramcao(args),
		exercicios_ninja_nivel_2.MenuExerciciosNinjaNivel2(args),
		fluxo_de_controle.MenuFluxoDeControle(args),
		exercicios_ninja_nivel_3.MenuExerciciosNinjaNivel3(args),
		agrupamento_de_dados.MenuAgrupamentoDeDados(args),
		exercicios_ninja_nivel_4.MenuExerciciosNinjaNivel4(args),
		structs.MenuStructs(args),
		exercicios_ninja_nivel_5.MenuExerciciosNinjaNivel5(args),
		funcoes.MenuFuncoes(args),
		exercicios_ninja_nivel_6.MenuExerciciosNinjaNivel6(args),
		ponteiros.MenuPonteiros(args),
		exercicios_ninja_nivel_7.MenuExercicioNinjaNivel7(args),
		aplicacoes.MenuAplicacoes(args),
		exercicios_ninja_nivel_8.MenuExerciciosNinjaNivel8(args),
		concorrencia.MenuConcorrencia(args),
		seu_ambiente_de_desenvolvimento.MenuSeuAmbienteDeDesenvolvimento(args),
		exercicios_ninja_nivel_9.MenuExerciciosNinjaNivel9(args),
		canais.MenuCanais(args),
		exercicios_ninja_nivel_10.MenuExerciciosNinjaNivel10(args),
		tratamento_de_erro.MenuTratamentoDeErro(args),
		exercicios_ninja_nivel_11.MenuExerciciosNinjaNivel11(args),
		documentacao.MenuDocumentacao(args),
		exercicios_ninja_nivel_12.MenuExerciciosNinjaNivel12(args),
	)
}

// generalOptions returns the general options available in the main menu.
// These options include help, outline, and chapter-specific options.
func generalOptions([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--help", ExecFunc: func() { HelpMe() }},
		{Options: "--outline", ExecFunc: func() { internal.Outline() }},
		{Options: "--caps", ExecFunc: func() { HelpMeCapituloOptions() }},
	}
}

// buildOptions combines the general and chapter-specific options based on the provided arguments.
// It formats and displays the menu options using the FormatMenuOptions function.
func buildOptions(args []string, options ...[]format.MenuOptions) {
	var opt []format.MenuOptions

	for _, o := range options {
		opt = append(opt, o...)
	}

	format.FormatMenuOptions(args, opt)
}
