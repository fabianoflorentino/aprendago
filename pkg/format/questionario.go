package format

import (
	"errors"
	"fmt"
	"strings"
)

// Interface CriaQuestionario para criar um questionário com as perguntas e opções de respostas
type CriaQuestionario interface {

	// O método Cria cria um questionário a partir de um slice do Questionario.
	// Retorna um slice de Questionario.
	Cria(Questionario []Questionario) []Questionario
}

// Interface ColetaRespostas para coletar as respostas de um questionário de uma prova
type ColetaRespostas interface {

	// O método Coleta coleta as respostas de um questionário de uma prova e retorna um slice de Resposta e um erro.
	// Recebe um slice de Questionario e retorna um slice de Resposta e um erro.
	Coleta(Questionario []Questionario) ([]Resposta, error)
}

// Interface ValidaResposta para validar as respostas de um questionário
type ValidaResposta interface {

	// O método Valida valida as respostas de um questionário e exibe o resultado da prova.
	// Recebe um slice de Resposta e um slice de Resposta.
	Valida(resposta []Resposta, gabarito []Resposta)
}

// Questionario de uma prova representa uma pergunta com suas opções de resposta.
type Questionario struct {
	Numero   string // Número da pergunta
	Pergunta string // Texto da pergunta
	Opcoes   string // Opções de resposta
}

// Resposta de uma pergunta armazena a resposta dada pelo usuário e seu status.
type Resposta struct {
	Numero   string // Número da pergunta
	Resposta string // Resposta dada pelo usuário
	Status   string // Opções de resposta disponíveis
}

// ServicoQuestionario é um tipo que implementa a interface CriaQuestionario através do método Cria
type ServicoQuestionario struct{}

// Metodo do tipo ServicoQuestionario que implementa a interface CriaQuestionario com o método Cria
// Recebe um slice de Questionario e retorna um slice de Questionario
func (sq ServicoQuestionario) Cria(questionario []Questionario) []Questionario {
	var prova []Questionario

	// Itera sobre o slice de Questionario e adiciona as perguntas e opções de respostas
	for _, quest := range questionario {
		prova = append(prova, Questionario{Numero: quest.Numero, Pergunta: quest.Pergunta, Opcoes: quest.Opcoes})
	}

	return prova
}

// ServicoResposta é um tipo que implementa a interface ColetaRespostas através do método Coleta
type ServicoResposta struct{}

// Metodo do tipo ServicoResposta que implementa a interface ColetaRespostas com o método Coleta
// Recebe um slice de Questionario e retorna um slice de Resposta e um erro
func (sr ServicoResposta) Coleta(questionario []Questionario) ([]Resposta, error) {
	var listaResposta []Resposta
	var resposta string

	// Cria um questionário com as perguntas e opções de respostas
	criar_questionario := new(ServicoQuestionario).Cria(questionario)

	// Itera sobre o slice de Questionario e coleta as respostas do usuário para cada pergunta
	for _, p := range criar_questionario {
		fmt.Print(p.Numero, p.Pergunta, p.Opcoes)
		fmt.Scan(&resposta)

		// Verifica se a resposta é válida
		if !validaOpcao(resposta, p.Opcoes) {
			return nil, errors.New("opção inválida")
		}

		// Adiciona a resposta ao slice de Resposta
		listaResposta = append(listaResposta, Resposta{p.Numero, resposta, ""})
	}

	return listaResposta, nil
}

// Metodo do tipo ServicoResposta que implementa a interface ValidaResposta com o método Valida
// Recebe um slice de Resposta e um slice de Resposta
func (sr ServicoResposta) Valida(resposta []Resposta, gabarito []Resposta) {
	var respostaValida []Resposta
	var respostaCorreta int

	// Itera sobre as respostas do usuário e verifica se a resposta é correta ou incorreta
	// Adiciona a resposta ao slice de RespostaValida
	for i, rsp := range resposta {
		if rsp.Resposta == gabarito[i].Resposta {
			respostaValida = append(respostaValida, Resposta{Numero: rsp.Numero, Resposta: rsp.Resposta, Status: "correta"})
			respostaCorreta++
		} else {
			respostaValida = append(respostaValida, Resposta{Numero: rsp.Numero, Resposta: rsp.Resposta, Status: "incorreta"})
		}
	}

	fmt.Printf("\nResultado da prova:\n\n")
	fmt.Printf("Total de respostas corretas %d de %d\n", respostaCorreta, len(gabarito))

	// Calcula a porcentagem de acertos do usuário na prova
	resultadoProva := float64(respostaCorreta) / float64(len(gabarito)) * 100
	fmt.Printf("Porcentagem de acertos %.2f%%\n\n", resultadoProva)

	// Exibe as respostas do usuário e o status de cada resposta
	for _, rsp := range respostaValida {

		// Remove o ponto final do número da pergunta
		rspNumero := strings.Trim(rsp.Numero, ".")

		fmt.Printf("Questão %s: Resposta: %s - %s\n", rspNumero, rsp.Resposta, rsp.Status)
	}
}

// validaOpcao verifica se a resposta do usuário é válida para uma pergunta.
// Recebe a resposta do usuário e as opções válidas como strings.
// Retorna true se a resposta estiver entre as opções válidas, caso contrário retorna false.
func validaOpcao(resposta, opcoes string) bool {
	// Itera sobre as opções de respostas da pergunta e verifica se a resposta do usuário é válida
	for _, opt := range strings.Fields(opcoes) {
		opt = strings.Trim(opt, "[]")
		opt = strings.Split(opt, " ")[0] // Extract the option letter
		if resposta == opt {
			return true
		}
	}
	return false
}
