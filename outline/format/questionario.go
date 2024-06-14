package format

import (
	"errors"
	"fmt"
	"strings"
)

// Questionario de uma prova representa uma pergunta com suas opções de resposta.
type TipoQuestionario struct {
	Numero   string // Número da pergunta
	Pergunta string // Texto da pergunta
	Opcoes   string // Opções de resposta
}

// Resposta de uma pergunta armazena a resposta dada pelo usuário e seu status.
type TipoResposta struct {
	Numero   string // Número da pergunta
	Resposta string // Resposta dada pelo usuário
	Status   string // Opções de resposta disponíveis
}

// ServicoQuestionario é um tipo que implementa a interface CriaQuestionario através do método Cria
type ServicoQuestionario struct{}

// ServicoResposta é um tipo que implementa a interface ColetaRespostas através do método Coleta
type ServicoResposta struct{}

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
func (sr *ServicoResposta) Valida(resposta []Resposta, gabarito []Resposta) {
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

	fmt.Printf("\nResultado da prova:\n")
	fmt.Printf("Total de respostas corretas %d de %d\n", respostaCorreta, len(gabarito))

	// Calcula a porcentagem de acertos do usuário na prova
	resultadoProva := float64(respostaCorreta) / float64(len(gabarito)) * 100
	fmt.Printf("Porcentagem de acertos %.2f%%\n", resultadoProva)

	for _, rsp := range respostaValida {
		fmt.Printf("Questão %s: Resposta %s - %s\n", rsp.Numero, rsp.Resposta, rsp.Status)
	}
}

// validaOpcao verifica se a resposta do usuário é válida para uma pergunta.
// Recebe a resposta do usuário e as opções válidas como strings.
// Retorna true se a resposta estiver entre as opções válidas, caso contrário retorna false.
func validaOpcao(resposta, opcoes string) bool {

	// Itera sobre as opções de respostas da pergunta e verifica se a resposta do usuário é válida
	for _, opt := range strings.Fields(opcoes) {
		opt = strings.Trim(opt, "[]")
		if resposta == opt {
			return true
		}
	}
	return false
}

/**
1. Interfaces

As interfaces permitem a definição de contratos que os tipos concretos devem seguir,
promovendo a abstração e a capacidade de intercambiar implementações.

CriaQuestionario: define um contrato para a criação de questionários.

- Método Cria: Permite criar um questionário a partir de uma lista de perguntas (Questionario).
Isso facilita a troca da implementação concreta sem alterar o código que usa essa interface.

ColetaRespostas: define um contrato para a coleta de respostas de um questionário.

- Método Coleta: Permite coletar as respostas dos usuários para um questionário. Através da interface,
podemos mudar como as respostas são coletadas (ex. input manual, leitura de um arquivo) sem
afetar o restante do código.

ValidaResposta: define um contrato para a validação das respostas de um questionário.

- Método Valida: Permite validar as respostas do usuário contra um gabarito fornecido.
A interface permite a alteração da lógica de validação (ex. diferentes critérios de correção)
sem impactar o restante do sistema.

2. Structs

As structs fornecem uma estrutura para armazenar dados relacionados a perguntas e respostas, além
de implementações concretas das interfaces.

Questionario: representa uma pergunta do questionário

Campos:
  * Numero: Armazena o número da pergunta
	* Pergunta: Armazena o texto da pergunta
	* Opcoes: opções disponíveis de resposta

Resposta: representa a resposta do usuário

Campos:
	* Numero: Armazena o número da pergunta
	* Resposta: Armazena a resposta dada pelo usuário
	* Status: Armazena o status da resposta (correta/incorreta)

ServicoQuestionario: implementa a interface CriaQuestionario

Método Cria: Responsável pela criação do questionário. Implementa o método da interface
CriaQuestionario que cria um questionário a partir de um slice de Questionario.

ServicoResposta: implementa a interface ColetaRespostas e ValidaResposta

Método Coleta: Responsável pela coleta das respostas do usuário. Implementa o método da interface
ColetaRespostas permitindo uma abordage flexível para coleta de respostas.

Método Valida: Responsável pela validação das respostas do usuário. Implementa o método da interface
ValidaResposta que valida as respostas do usuário contra um gabarito fornecido.

3. Funções
As funções realizam operações específicas que são necessárias para o fluxo de trabalho do questionário, tais como validação de opções e coleta de respostas.

opcaoValida

Verifica se uma resposta está entre as opções válidas para uma pergunta.

Parâmetros resposta, opcoes: Recebe a resposta fornecida pelo usuário e as opções válidas.
Retorna bool: Retorna true se a resposta está entre as opções válidas, false caso contrário. Facilita a validação de respostas.


Exemplo de Uso: Cria instâncias dos serviços (QuestionnaireService e AnswerService),
cria um questionário, coleta respostas e valida essas respostas. Demonstra como as
interfaces permitem a substituição e a flexibilidade das implementações concretas.

---
Benefícios das Interfaces e Estruturas na Refatoração

Modularidade:

As interfaces separam o contrato da implementação, permitindo a troca de uma
implementação por outra sem alterar o código dependente.

Testabilidade:

Facilita a criação de mocks e stubs para testes unitários, permitindo testar
componentes de forma isolada.

Manutenibilidade:

O código é mais fácil de manter, pois a alteração de uma implementação concreta não
afeta a interface ou outras implementações que seguem o mesmo contrato.

Flexibilidade:

Permite a adição de novas formas de coleta, criação e validação sem mudar a estrutura
principal do código. Por exemplo, podemos adicionar um novo AnswerCollector que coleta
respostas de uma API em vez de fmt.Scan.
---
**/
