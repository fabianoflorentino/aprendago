package exercicios_ninja_nivel_1

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func ResolucaoNaPraticaExercicio1() {
	x := 42
	y := "James Bond"
	z := true

	resolucao := fmt.Sprintf("%v %v %v", x, y, z)

	format.FormatResolucaoExercicios(resolucao)
}

func ResolucaoNaPraticaExercicio2() {
	var x int
	var y string
	var z bool

	resolucao := fmt.Sprintf("%v %v %v", x, y, z)

	format.FormatResolucaoExercicios(resolucao)
}

func ResolucaoNaPraticaExercicio3() {
	x := 42
	y := "James Bond"
	z := true

	resolucao := fmt.Sprintf("%v %v %v", x, y, z)

	format.FormatResolucaoExercicios(resolucao)
}

func ResolucaoNaPraticaExercicio4() {
	type ninja int

	x := ninja(42)

	resolucao := fmt.Sprintf("%v", x)
	format.FormatResolucaoExercicios(resolucao)

	fmt.Printf("\n%T\n", x)

}

func ResolucaoNaPraticaExercicio5() {
	type ninja int

	var x ninja
	var y int

	x = 42
	y = int(x)

	resolucao := fmt.Sprintf("%T\n%T", x, y)

	format.FormatResolucaoExercicios(resolucao)
}

func RespondaAProva() {
	questionario := []format.Questionnaire{
		{Number: "1.", Question: "Qual o menor elemento em um programa que expressa uma ação a ser executada?", Options: "[1] Uma declaração (Statement) [2] Uma expressão: "},
		{Number: "2.", Question: "A combinação de um ou mais valores, constantes, variáveis, operadores e funções que a linguagem interpreta e usa para produzir outro valor é?", Options: "[1] Uma declaração (Statement) [2] Uma expressão: "},
		{Number: "3.", Question: "Quais são parênteses?", Options: "[1] () [2] {} [3] []: "},
		{Number: "4.", Question: "Quais são colchetes?", Options: "[1] () [2] {} [3] []: "},
		{Number: "5.", Question: "Quais são chaves?", Options: "[1] () [2] {} [3] []: "},
		{Number: "6.", Question: "A abrangência de uma variável designa onde no código você pode acessar essa variável, e atribuir ou ler valores dela", Options: "[1] Verdade [2] Mentira: "},
		{Number: "7.", Question: "Um tipo de dados primitivo é um tipo composto, criado a partir de outros tipos básicos que ja vem de fabrica na linguagem", Options: "[1] Verdade [2] Mentira: "},
		{Number: "8.", Question: "O tipo int é um tipo de dado primitivo", Options: "[1] Verdade [2] Mentira: "},
		{Number: "9.", Question: "O tipo string é um tipo de dado composto", Options: "[1] Verdade [2] Mentira: "},
		{Number: "10.", Question: "Um tipo de dado composto permite que você crie estruturas formadas de outros tipos de dados", Options: "[1] Verdade [2] Mentira: "},
		{Number: "11.", Question: "Quando declaramos uma variável com a palavra chave var e não atribuimos nenhum valor a esta variável, o compilador designa para esta um valor padrão, chamado de valor zero.", Options: "[1] Verdade [2] Mentira: "},
		{Number: "12.", Question: "Palavras chaves servem a propositos especificos, mas fora isso podem ser usadas livremente ao longo do programa", Options: "[1] Verdade [2] Mentira: "},
		{Number: "13.", Question: "Palavra chave e palavra reservada designam coisas diferentes", Options: "[1] Verdade [2] Mentira: "},
		{Number: "14.", Question: "Uma palavra chave somente pode ser usada para seu proposito especifico e para declarar variaveis", Options: "[1] Verdade [2] Mentira: "},
		{Number: "15.", Question: "Em 2 + 2, os numeros 2 são operadores", Options: "[1] Verdade [2] Mentira: "},
		{Number: "16.", Question: "O termo package é uma palavra chave", Options: "[1] Verdade [2] Mentira: "},
		{Number: "17.", Question: "O termo variable é uma palavra chave", Options: "[1] Verdade [2] Mentira: "},
		{Number: "18.", Question: "O ponto de entrada para todos os programas é a função main(), que deve ficar dentro do package main", Options: "[1] Verdade [2] Mentira: "},
		{Number: "19.", Question: "O operador curto de declaração pode ser usado ao invés de var em todas as situações", Options: "[1] Verdade [2] Mentira: "},
		{Number: "20.", Question: "Quando vemos fmt.Println(), isto esta chamando a função Println() que pertence ao package fmt", Options: "[1] Verdade [2] Mentira: "},
		{Number: "21.", Question: "Identificador é o nome atribuido a uma variável, função ou constante", Options: "[1] Verdade [2] Mentira: "},
		{Number: "22.", Question: "Para utilizar uma função, variável ou constante de outro package utiliza-se o formato package-ponto-identificador. Por exemplo fmt.Println()", Options: "[1] Verdade [2] Mentira: "},
		{Number: "23.", Question: "Qual é o caracter que permite jogar fora um valor? Ou seja, qual caracter permite que voce diga ao programa que não vai utilzar o valor retornado por uma função?", Options: "[1] # [2] @ [3] _ [4] $ [5] isso é pegadinha...: "},
		{Number: "24.", Question: "Uma função cujo parâmetro é '... interface{}' é uma função variatica. Isso significa que você pode passar à função um número pré determinado de valores", Options: "[1] Verdade [2] Mentira: "},
		{Number: "25.", Question: "Todo valor em Go pertence também ao tipo interface vazia, representado pela notação interface{}", Options: "[1] Verdade [2] Mentira: "},
		{Number: "26.", Question: "2+3 é uma declaração/statement, não uma expressão", Options: "[1] Verdade [2] Mentira: "},
		{Number: "27.", Question: "2+3 é uma expressão não uma declaração/statemente", Options: "[1] Verdade [2] Mentira: "},
		{Number: "28.", Question: "Se eu quiser salvar o resultado de um format printing em uma variável, posso usar a função fmt.Sprintf()", Options: "[1] Verdade [2] Mentira: "},
		{Number: "29.", Question: "Em Go podemos criar nossos proprios tipos?", Options: "[1] Verdade [2] Mentira: "},
		{Number: "30.", Question: "Falando de tipos, em Go utilizamos o termo 'coerção' diferentemente de Java, por exemplo, onde se utiliza o termo 'conversão'", Options: "[1] Verdade [2] Mentira: "},
		{Number: "31.", Question: "Todo tipo criado pelo programador tem um tipo subjacente", Options: "[1] Verdade [2] Mentira: "},
	}

	gabarito := []format.Answer{
		{Number: "1.", Answer: "1"}, {Number: "2.", Answer: "2"}, {Number: "3.", Answer: "1"}, {Number: "4.", Answer: "3"}, {Number: "5.", Answer: "2"},
		{Number: "6.", Answer: "1"}, {Number: "7.", Answer: "2"}, {Number: "8.", Answer: "1"}, {Number: "9.", Answer: "2"}, {Number: "10.", Answer: "1"},
		{Number: "11.", Answer: "1"}, {Number: "12.", Answer: "2"}, {Number: "13.", Answer: "1"}, {Number: "14.", Answer: "2"}, {Number: "15.", Answer: "2"},
		{Number: "16.", Answer: "1"}, {Number: "17.", Answer: "2"}, {Number: "18.", Answer: "1"}, {Number: "19.", Answer: "2"}, {Number: "20.", Answer: "1"},
		{Number: "21.", Answer: "1"}, {Number: "22.", Answer: "1"}, {Number: "23.", Answer: "3"}, {Number: "24.", Answer: "1"}, {Number: "25.", Answer: "1"},
		{Number: "26.", Answer: "2"}, {Number: "27.", Answer: "1"}, {Number: "28.", Answer: "1"}, {Number: "29.", Answer: "1"}, {Number: "30.", Answer: "1"}, {Number: "31.", Answer: "1"},
	}

	// Cria o questionário
	sq := new(format.QuestionnaireService)
	sr := new(format.AnswerService)

	prova := sq.Create(questionario)
	respostas, err := sr.Collect(prova)
	if err != nil {
		fmt.Println("Erro ao coletar respostas:", err)
	}

	sr.Validate(respostas, gabarito)
}
