/*
Package menu fornece uma interface de linha de comando para navegar pelas opções e funcionalidades
do curso Aprenda GO. Ele agrupa e organiza diferentes menus e suas opções, permitindo ao usuário
acessar capítulos específicos, exercícios práticos e visões gerais do curso com base nos argumentos fornecidos.

Funcionalidades principais:

- Exibe o menu principal com opções gerais e específicas dos capítulos.
- Permite acessar diferentes seções do curso, como variáveis, tipos, fluxo de controle e exercícios práticos.
- Processa os argumentos fornecidos pelo usuário para determinar e executar a opção apropriada.
- Oferece suporte a opções de ajuda e visão geral para facilitar a navegação.

Uso básico:

Para usar o pacote menu, você deve chamar a função Options com uma lista de argumentos.
Esses argumentos determinam qual menu ou opção será exibido e executado.

Exemplo:

	args := []string{"--help"}
	menu.Options(args)

Nesse exemplo, a função Options exibirá a ajuda, com base no argumento fornecido.
*/
package menu

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/internal"
	"github.com/fabianoflorentino/aprendago/internal/agrupamento_de_dados"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_1"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_2"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_3"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_4"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_5"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_6"
	"github.com/fabianoflorentino/aprendago/internal/fluxo_de_controle"
	"github.com/fabianoflorentino/aprendago/internal/funcoes"
	"github.com/fabianoflorentino/aprendago/internal/fundamentos_da_programacao"
	"github.com/fabianoflorentino/aprendago/internal/structs"
	"github.com/fabianoflorentino/aprendago/internal/variaveis_valores_tipos"
	"github.com/fabianoflorentino/aprendago/internal/visao_geral_do_curso"
	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// Options é a função principal do pacote menu. Ela exibe o menu principal do programa,
// mostrando todas as opções disponíveis com base nos argumentos fornecidos (args).
func Options(args []string) {
	fmt.Printf("Aprenda GO\n\n")

	// Chama a função buildOptions para construir e exibir as opções de menu.
	// Combina opções gerais e específicas dos capítulos com base nos argumentos fornecidos.
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
	)
}

// generalOptions retorna as opções gerais disponíveis em todos os menus,
// como ajuda e visão geral do curso.
func generalOptions([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--help", ExecFunc: func() { HelpMe() }},
		{Options: "--outline", ExecFunc: func() { internal.Outline() }},
		{Options: "--caps", ExecFunc: func() { HelpMeCapituloOptions() }},
	}
}

// buildOptions combina e processa as opções de menu com base nos argumentos fornecidos.
// Ele verifica se os argumentos correspondem a alguma opção de menu e executa a ação associada.
func buildOptions(args []string, options ...[]format.MenuOptions) {
	var opt []format.MenuOptions

	// Combina todas as opções recebidas em uma única lista.
	for _, o := range options {
		opt = append(opt, o...)
	}

	// Formata e exibe as opções de menu com base nos argumentos.
	format.FormatMenuOptions(args, opt)
}
