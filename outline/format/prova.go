// Pacote com funções para criar um questionário, coletar respostas e validar as respostas
package format

import (
	"errors"
	"fmt"
	"strings"
)

// Questionario de uma prova
type Questionario struct {
	Numero   string
	Pergunta string
	Opcoes   string
}

// Resposta de uma pergunta
type Resposta struct {
	Numero   string
	Resposta string
	Status   string
}

// CriarQuestionario cria um questionário com as perguntas e opções de respostas de uma prova
// Recebe um slice de Questionario e retorna um slice de Questionario
func CriarQuestionario(questionario []Questionario) []Questionario {
	var prova []Questionario

	// Itera sobre o slice de Questionario e adiciona as perguntas e opções de respostas
	for _, quest := range questionario {
		prova = append(prova, Questionario{Numero: quest.Numero, Pergunta: quest.Pergunta, Opcoes: quest.Opcoes})
	}

	return prova
}

// ColetarRespostas coleta as respostas de um questionário de uma prova e retorna um slice de Resposta
// Recebe um slice de Questionario e retorna um slice de Resposta e um erro
func ColetarRespostas(questionario []Questionario) ([]Resposta, error) {
	var listaRespostas []Resposta
	var resposta string

	criar_questionario := CriarQuestionario(questionario)

	// Itera sobre o slice de Questionario e coleta as respostas do usuário para cada pergunta
	// Verifica se a resposta é válida e adiciona ao slice de Resposta ou retorna um erro
	for _, p := range criar_questionario {
		fmt.Print(p.Numero, p.Pergunta, p.Opcoes)
		fmt.Scan(&resposta)

		// Verifica se a resposta é válida
		if !opcaoValida(resposta, p.Opcoes) {
			return nil, errors.New("opção inválida")
		}

		listaRespostas = append(listaRespostas, Resposta{p.Numero, resposta, ""})
	}

	return listaRespostas, nil
}

// opcaoValida verifica se a resposta do usuário é válida para a pergunta da prova e retorna um boolean
// Recebe a resposta e as opções de respostas da pergunta e retorna um boolean
func opcaoValida(resposta, opcoes string) bool {
	// Itera sobre as opções de respostas da pergunta e verifica se a resposta do usuário é válida
	// Retorna true se a resposta for válida e false se for inválida
	for _, opt := range strings.Fields(opcoes) {
		opt = strings.Trim(opt, "[]")
		if resposta == opt {
			return true
		}
	}

	return false
}

// ValidarRespostas valida as respostas do usuário com o gabarito da prova e exibe o resultado da prova na tela
// Recebe um slice de Resposta e um slice de Resposta
func ValidarRespostas(respostas []Resposta, gabarito []Resposta) {
	var validaRespostas []Resposta
	var totalCorretas int

	// Itera sobre as respostas do usuário e compara com o gabarito da prova
	for i, resp := range respostas {
		if resp.Resposta == gabarito[i].Resposta {
			validaRespostas = append(validaRespostas, Resposta{resp.Numero, resp.Resposta, "Correta"})
			totalCorretas++
		} else {
			validaRespostas = append(validaRespostas, Resposta{resp.Numero, resp.Resposta, "Incorreta"})
		}
	}

	// Exibe o resultado da prova na tela
	fmt.Printf("\nResultado da prova:\n")
	fmt.Printf("\nTotal de respostas corretas: %d de %d\n", totalCorretas, len(gabarito))

	// Calcula o percentual de acerto do usuário na prova
	resultadoDaProva := float64(totalCorretas) / float64(len(gabarito)) * 100
	fmt.Printf("Percentual de acerto: %.2f%%.\n\n", resultadoDaProva)

	// Exibe as respostas do usuário e o status de cada resposta
	for _, v := range validaRespostas {
		fmt.Printf("P: %-4s R: %-4s %s\n", v.Numero, v.Resposta, v.Status)
	}
}
