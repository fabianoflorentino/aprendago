// Prova para validar o conhecimento dos primeiros capitulos do curso
package outline

import (
	"errors"
	"fmt"
	"strings"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

// Prova representa uma pergunta da prova
type Prova struct {
	NumeroPergunta string
	Perguntas      string
	Opcoes         string
}

// Resposta representa uma resposta fornecia pelo usuário
type Resposta struct {
	NumeroPergunta string
	Resposta       string
}

// NaPraticaExercicio6 é a função que exibe o enunciado do exercício 6
func NaPraticaExercicio6() {
	topic := format.OutlineContent{
		Title: "Na Prática - Exercício #6",
		Content: `
- Prova!
- Link: https://goo.gl/forms/s9y91iVSPvA4iahj1 [Link quebrado]. (Use a opção --na-pratica-exercicio-6 --nivel-1 --prova)
- Se você deu pausa e fez todos os exercícios anteriores você mesmo, e só viu a resposta depois...
  e se você der pausa agora e fizer a prova inteira por conta própria, e só assistir as respostas depois...
  sabe o que isso quer dizer? Que você é ninja. Ninja nível 1. Tá no caminho certo pra ser um developer ninja mestre.
    `,
	}

	format.FormatOutlineTopic(topic)
}

// CriarQuestionario retorna uma lista de perguntas da prova
func CriarQuestionario() []Prova {
	return []Prova{
		{"1.", "Qual o menor elemento em um programa que expressa uma ação a ser executada?", "[1] Uma declaração (Statement) [2] Uma expressão: "},
		{"2.", "A combinação de um ou mais valores, constantes, variáveis, operadores e funções que a linguagem interpreta e usa para produzir outro valor é?", "[1] Uma declaração (Statement) [2] Uma expressão: "},
		{"3.", "Quais são parênteses?", "[1] () [2] {} [3] []: "},
		{"4.", "Quais são colchetes?", "[1] () [2] {} [3] []: "},
		{"5.", "Quais são chaves?", "[1] () [2] {} [3] []: "},
		{"6.", "A abrangência de uma variável designa onde no código você pode acessar essa variável, e atribuir ou ler valores dela", "[1] Verdade [2] Mentira: "},
		{"7.", "Um tipo de dados primitivo é um tipo composto, criado a partir de outros tipos básicos que ja vem de fabrica na linguagem", "[1] Verdade [2] Mentira: "},
		{"8.", "O tipo int é um tipo de dado primitivo", "[1] Verdade [2] Mentira: "},
		{"9.", "O tipo string é um tipo de dado composto", "[1] Verdade [2] Mentira: "},
		{"10.", "Um tipo de dado composto permite que você crie estruturas formadas de outros tipos de dados", "[1] Verdade [2] Mentira: "},
		{"11.", "Quando declaramos uma variável com a palavra chave var e não atribuimos nenhum valor a esta variável, o compilador designa para esta um valor padrão, chamado de valor zero.", "[1] Verdade [2] Mentira: "},
		{"12.", "Palavras chaves servem a propositos especificos, mas fora isso podem ser usadas livremente ao longo do programa", "[1] Verdade [2] Mentira: "},
		{"13.", "Palavra chave e palavra reservada designam coisas diferentes", "[1] Verdade [2] Mentira: "},
		{"14.", "Uma palavra chave somente pode ser usada para seu proposito especifico e para declarar variaveis", "[1] Verdade [2] Mentira: "},
		{"15.", "Em 2 + 2, os numeros 2 são operadores", "[1] Verdade [2] Mentira: "},
		{"16.", "O termo package é uma palavra chave", "[1] Verdade [2] Mentira: "},
		{"17.", "O termo variable é uma palavra chave", "[1] Verdade [2] Mentira: "},
		{"18.", "O ponto de entrada para todos os programas é a função main(), que deve ficar dentro do package main", "[1] Verdade [2] Mentira: "},
		{"19.", "O operador curto de declaração pode ser usado ao invés de var em todas as situações", "[1] Verdade [2] Mentira: "},
		{"20.", "Quando vemos fmt.Println(), isto esta chamando a função Println() que pertence ao package fmt", "[1] Verdade [2] Mentira: "},
		{"21.", "Identificador é o nome atribuido a uma variável, função ou constante", "[1] Verdade [2] Mentira: "},
		{"22.", "Para utilizar uma função, variável ou constante de outro package utiliza-se o formato package-ponto-identificador. Por exemplo fmt.Println()", "[1] Verdade [2] Mentira: "},
		{"23.", "Qual é o caracter que permite jogar fora um valor? Ou seja, qual caracter permite que voce diga ao programa que não vai utilzar o valor retornado por uma função?", "[1] # [2] @ [3] _ [4] $ [5] isso é pegadinha...: "},
		{"24.", "Uma função cujo parâmetro é '... interface{}' é uma função variatica. Isso significa que você pode passar à função um número pré determinado de valores", "[1] Verdade [2] Mentira: "},
		{"25.", "Todo valor em Go pertence também ao tipo interface vazia, representado pela notação interface{}", "[1] Verdade [2] Mentira: "},
		{"26.", "2+3 é uma declaração/statement, não uma expressão", "[1] Verdade [2] Mentira: "},
		{"27.", "2+3 é uma expressão não uma declaração/statemente", "[1] Verdade [2] Mentira: "},
		{"28.", "Se eu quiser salvar o resultado de um format printing em uma variável, posso usar a função fmt.Sprintf()", "[1] Verdade [2] Mentira: "},
		{"29.", "Em Go podemos criar nossos proprios tipos?", "[1] Verdade [2] Mentira: "},
		{"30.", "Falando de tipos, em Go utilizamos o termo 'coerção' diferentemente de Java, por exemplo, onde se utiliza o termo 'conversão'", "[1] Verdade [2] Mentira: "},
		{"31.", "Todo tipo criado pelo programador tem um tipo subjacente", "[1] Verdade [2] Mentira: "},
	}
}

// ColetarRespostas coleta as respostas do usuário para as perguntas da prova
func ColetarRespostas(perguntas []Prova) ([]Resposta, error) {
	var listaRespostas []Resposta
	var resposta string

	// Exibe as perguntas da prova e coleta as respostas do usuário
	for _, pergunta := range perguntas {
		fmt.Println(pergunta.NumeroPergunta, pergunta.Perguntas)
		fmt.Println(pergunta.Opcoes)
		fmt.Scan(&resposta)

		if !opcaoValida(resposta, pergunta.Opcoes) {
			return nil, errors.New("opção inválida")
		}

		// Adiciona as respostas do usuário a lista de respostas concatenando o número da pergunta e a resposta
		// gerando uma lista
		listaRespostas = append(listaRespostas, Resposta{pergunta.NumeroPergunta, resposta})
	}

	return listaRespostas, nil
}

// opcaoValida verifica se a opção fornecida pelo usuário
func opcaoValida(resposta, opcoes string) bool {
	for _, opt := range strings.Fields(opcoes) {
		opt = strings.Trim(opt, "[]")
		if resposta == opt {
			return true
		}
	}

	return false
}

// ImprimirRespostas exibe as respostas fornecidas pelo usuário
func ImprimirRespostas(respostas []Resposta) {
	fmt.Println("Respostas da prova:")

	// Exibe as respostas fornecidas pelo usuário em duas colunas de 15 perguntas com suas respectivas respostas
	for i := 0; i < len(respostas); i++ {
		if i+1 < len(respostas) && i+15 < len(respostas) {
			fmt.Printf("P: %-4s R: %-8s P: %-4s R: %-8s\n", respostas[i].NumeroPergunta, respostas[i].Resposta, respostas[i+15].NumeroPergunta, respostas[i+15].Resposta)
		}
	}
}

// ImprimiGabarito exibe o gabarito da prova
func ImprimiGabarito() {
	gabarito := []Resposta{
		{"1.", "1"}, {"2.", "2"}, {"3.", "1"}, {"4.", "3"}, {"5.", "2"}, {"6.", "1"}, {"7.", "2"}, {"8.", "1"}, {"9.", "2"}, {"10.", "1"}, {"11.", "1"},
		{"12.", "2"}, {"13.", "1"}, {"14.", "2"}, {"15.", "2"}, {"16.", "1"}, {"17.", "2"}, {"18.", "1"}, {"19.", "2"}, {"20.", "1"}, {"21.", "1"},
		{"22.", "1"}, {"23.", "3"}, {"24.", "1"}, {"25.", "1"}, {"26.", "2"}, {"27.", "1"}, {"28.", "1"}, {"29.", "1"}, {"30.", "1"}, {"31.", "1"},
	}

	// Exibe o gabarito da prova para o usuário em  duas colunas de 15 perguntas com suas respectivas respostas
	fmt.Println("Gabarito da prova:")
	for i := 0; i < len(gabarito); i++ {
		if i+1 < len(gabarito) && i+15 < len(gabarito) {
			fmt.Printf("P: %-4s R: %-8s P: %-4s R: %-8s\n", gabarito[i].NumeroPergunta, gabarito[i].Resposta, gabarito[i+15].NumeroPergunta, gabarito[i+15].Resposta)
		}
	}
}

func RespondaAProva() {
	perguntas := CriarQuestionario()
	respostas, err := ColetarRespostas(perguntas)
	if err != nil {
		fmt.Println("Erro ao coletar respostas:", err)
		return
	}

	ImprimirRespostas(respostas)
	ImprimiGabarito()
}
