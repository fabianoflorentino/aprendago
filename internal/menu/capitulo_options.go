package menu

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fabianoflorentino/aprendago/internal/agrupamento_de_dados"
	"github.com/fabianoflorentino/aprendago/internal/aplicacoes"
	"github.com/fabianoflorentino/aprendago/internal/canais"
	"github.com/fabianoflorentino/aprendago/internal/concorrencia"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_1"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_10"
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
	"github.com/fabianoflorentino/aprendago/pkg/format"
	"github.com/fabianoflorentino/aprendago/pkg/reader"
)

const MAX_CHAPTER = 22

func MenuCapituloOptions([]string) []format.MenuOptions {
	menuOptions := []format.MenuOptions{}

	for chapterNumber := 1; chapterNumber <= MAX_CHAPTER; chapterNumber++ {
		menuOption := format.MenuOptions{
			Options:  "--cap=" + strconv.Itoa(chapterNumber) + " --topics",
			ExecFunc: menuOptionFuncToExecute(chapterNumber),
		}

		menuOptions = append(menuOptions, menuOption)
	}

	return menuOptions
}

func HelpMeCapituloOptions() {
	helpMes := []format.HelpMe{}

	for chapterNumber := 1; chapterNumber <= MAX_CHAPTER; chapterNumber++ {
		helpMe := format.HelpMe{
			Flag:        "--cap=" + string(chapterNumber) + " --topics",
			Description: helpMeChapterDescription(chapterNumber),
			Width:       0,
		}

		helpMes = append(helpMes, helpMe)
	}

	fmt.Println("Capítulos do Curso")
	format.PrintHelpMe(helpMes)
}

// Função migratória, para cada capítulo que for criado no internal/outlines, o case pode ser removido
// quando não restar mais nenhum case, o conteúdo da clásula default pode ser movido para a função MenuCapituloOptions
func menuOptionFuncToExecute(chapterNumber int) func() {
	switch chapterNumber {
	case 3:
		return func() { exercicios_ninja_nivel_1.HelpMeExerciciosNinjaNivel1() }
	case 4:
		return func() { fundamentos_da_programacao.HelpMeFundamentosDaProgramacao() }
	case 5:
		return func() { exercicios_ninja_nivel_2.HelpMeExerciciosNinjaNivel2() }
	case 6:
		return func() { fluxo_de_controle.HelpMeFluxoDeControle() }
	case 7:
		return func() { exercicios_ninja_nivel_3.HelpMeExerciciosNinjaNivel3() }
	case 8:
		return func() { agrupamento_de_dados.HelpMeAgrupamentoDeDados() }
	case 9:
		return func() { exercicios_ninja_nivel_4.HelpMeExerciciosNinjaNivel4() }
	case 10:
		return func() { structs.HelpMeStructs() }
	case 11:
		return func() { exercicios_ninja_nivel_5.HelpMeExerciciosNinjaNivel5() }
	case 12:
		return func() { funcoes.HelpMeFuncoes() }
	case 13:
		return func() { exercicios_ninja_nivel_6.HelMeExerciciosNinjaNivel6() }
	case 14:
		return func() { ponteiros.HelpMePonteiros() }
	case 15:
		return func() { exercicios_ninja_nivel_7.HelpMeExerciciosNinjaNivel7() }
	case 16:
		return func() { aplicacoes.HelpMeAplicacoes() }
	case 17:
		return func() { exercicios_ninja_nivel_8.HelpMeExerciciosNinjaNivel8() }
	case 18:
		return func() { concorrencia.HelpMeConcorrencia() }
	case 19:
		return func() { seu_ambiente_de_desenvolvimento.HelpMeSeuAmbienteDeDesenvolvimento() }
	case 20:
		return func() { exercicios_ninja_nivel_9.HelpMeExerciciosNinjaNivel9() }
	case 21:
		return func() { canais.HelpMeCanais() }
	case 22:
		return func() { exercicios_ninja_nivel_10.HelpMeExerciciosNinjaNivel10() }
	default:
		return func() { format.PrintChapterHelpMe(chapterNumber) }
	}
}

// Função migratória, para cada capítulo que for criado no internal/outlines, o case pode ser removido
// quando não restar mais nenhum case, o conteúdo da clásula default pode ser movido para a função HelpMeCapituloOptions
func helpMeChapterDescription(chapterNumber int) string {
	switch chapterNumber {
	case 3:
		return "Exercícios Ninja: Nível 1"
	case 4:
		return "Fundamentos da Programação"
	case 5:
		return "Exercícios Ninja: Nível 2"
	case 6:
		return "Fluxo de Controle"
	case 7:
		return "Exercícios Ninja: Nível 3"
	case 8:
		return "Agrupamento de Dados"
	case 9:
		return "Exercícios Ninja: Nível 4"
	case 10:
		return "Structs"
	case 11:
		return "Exercícios Ninja: Nível 5"
	case 12:
		return "Funções"
	case 13:
		return "Exercícios Ninja: Nível 6"
	case 14:
		return "Ponteiros"
	case 15:
		return "Exercícios Ninja: Nível 7"
	case 16:
		return "Aplicações"
	case 17:
		return "Exercícios Ninja: Nível 8"
	case 18:
		return "Concorrência"
	case 19:
		return "Seu Ambiente de Desenvolvimento"
	case 20:
		return "Exercícios Ninja: Nível 9"
	case 21:
		return "Canais"
	case 22:
		return "Exercícios Ninja: Nível 10"
	default:
		chapter, error := reader.ReadChapterFile(chapterNumber)
		if error != nil {
			log.Panicf("Erro ai tentar ler arquivo do Capítulo %d", chapterNumber)
		}
		return chapter.Title
	}
}
