// Package exercicios_ninja_nivel_2 provides solutions to various exercises
// and a questionnaire related to Go programming language concepts. The
// exercises cover topics such as constants, bitwise operations, and string
// formatting. The questionnaire includes multiple-choice questions to test
// knowledge on Go language fundamentals.
package exercicios_ninja_nivel_2

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

// ResolucaoNaPraticaExercicio1 prints the number 1000019 in decimal, binary, and hexadecimal formats.
// It formats the number using fmt.Sprintf and then passes the formatted string to the
// FormatResolucaoExercicios function from the format package.
func ResolucaoNaPraticaExercicio1() {
	numero := 1000019

	resoluca := fmt.Sprintf("Decimal: %d \nBinário: %b \nHexadecimal: %#x \n", numero, numero, numero)

	format.FormatResolucaoExercicios(resoluca)
}

// ResolucaoNaPraticaExercicio2 demonstrates the use of comparison operators in Go.
// It initializes two integer variables, `a` and `b`, with values 10 and 20 respectively.
// The function then creates a formatted string `resolucao` that includes the values of `a` and `b`,
// as well as the results of various comparison operations between `a` and `b`:
// - a == b: checks if `a` is equal to `b`
// - a != b: checks if `a` is not equal to `b`
// - a <= b: checks if `a` is less than or equal to `b`
// - a < b: checks if `a` is less than `b`
// - a >= b: checks if `a` is greater than or equal to `b`
// - a > b: checks if `a` is greater than `b`
// Finally, it calls the `FormatResolucaoExercicios` function from the `format` package to format and display the results.
func ResolucaoNaPraticaExercicio2() {
	a := 10
	b := 20

	resolucao := fmt.Sprintf("a: %v \nb: %v \n\na == b: %v \na != b: %v \na <= b: %v \na < b: %v \na >= b: %v \na > b: %v \n", a, b, a == b, a != b, a <= b, a < b, a >= b, a > b)

	format.FormatResolucaoExercicios(resolucao)
}

// ResolucaoNaPraticaExercicio3 demonstrates the use of both untyped and typed constants in Go.
// It defines two constants: `untypedConst` with an untyped value of 10, and `typedConst` with a typed value of 20.
// The function then formats these constants into a string using `fmt.Sprintf` and passes the formatted string
// to the `FormatResolucaoExercicios` function from the `format` package for further processing or display.
func ResolucaoNaPraticaExercicio3() {
	const (
		untypedConst     = 10
		typedConst   int = 20
	)

	resolucao := fmt.Sprintf("untypedConst: %v \ntypedConst: %v", untypedConst, typedConst)

	format.FormatResolucaoExercicios(resolucao)
}

// ResolucaoNaPraticaExercicio4 demonstrates the conversion of an integer value
// to its decimal, binary, and hexadecimal representations. It performs the following steps:
// 1. Initializes an integer variable v1 with the value 10.
// 2. Formats the value of v1 into a string showing its decimal, binary, and hexadecimal forms.
// 3. Left shifts the value of v1 by 1 bit to obtain v2.
// 4. Formats the value of v2 into a string showing its decimal, binary, and hexadecimal forms.
// 5. Concatenates the formatted strings of v1 and v2 and passes the result to the FormatResolucaoExercicios function for further processing or display.
func ResolucaoNaPraticaExercicio4() {
	v1 := 10

	resolucao1 := fmt.Sprintf("Decimal: %d \nBinário: %b \nHexadecimal: %#x \n\n", v1, v1, v1)

	v2 := v1 << 1

	resolucao2 := fmt.Sprintf("Decimal: %d \nBinário: %b \nHexadecimal: %#x", v2, v2, v2)

	format.FormatResolucaoExercicios(resolucao1 + resolucao2)
}

// ResolucaoNaPraticaExercicio5 demonstrates a simple function that assigns a string
// to a variable and then formats it using the FormatResolucaoExercicios function
// from the format package. The string assigned is "Aprenda Go!" which means "Learn Go!" in Portuguese.
func ResolucaoNaPraticaExercicio5() {
	v1 := `Aprenda Go!`

	format.FormatResolucaoExercicios(v1)
}

// ResolucaoNaPraticaExercicio6 demonstrates the use of the iota identifier in Go.
// It defines a constant block where the first constant 'ano' is set to 2024, and
// subsequent constants 'v1', 'v2', 'v3', and 'v4' are automatically incremented by 1
// starting from 2025. The function then formats these values into a string and
// passes it to the FormatResolucaoExercicios function for further processing.
func ResolucaoNaPraticaExercicio6() {
	const (
		ano = 2024 + iota
		v1
		v2
		v3
		v4
	)

	resolucao := fmt.Sprintf("v1: %v\nv2: %v\nv3: %v\nv4: %v\n", v1, v2, v3, v4)

	format.FormatResolucaoExercicios(resolucao)
}

// RespondaAProva is a function that simulates a quiz with multiple-choice questions.
// It initializes a set of questions and their corresponding correct answers, creates a questionnaire,
// collects responses from the user, and validates the responses against the correct answers.
//
// The function performs the following steps:
// 1. Initializes a slice of questions (questionario) with their respective options.
// 2. Initializes a slice of correct answers (gabarito) corresponding to the questions.
// 3. Creates instances of QuestionnaireService and AnswerService.
// 4. Uses the QuestionnaireService to create a questionnaire from the list of questions.
// 5. Uses the AnswerService to collect responses from the user for the created questionnaire.
// 6. If there is an error during the collection of responses, it prints an error message.
// 7. Validates the collected responses against the correct answers using the AnswerService.
func RespondaAProva() {
	questionario := []format.Questionnaire{
		{Number: "1.", Question: "Ao utilizar o opeardor curto de declaração, não é necessário especificar o tipo", Options: "[1] Verdadeiro [2] Falso: "},
		{Number: "2.", Question: "Uma variável do tipo bool pode conter três valores: true, false, undefined", Options: "[1] Verdadeiro [2] Falso: "},
		{Number: "3.", Question: "Se você tiver 5 lâmpadas na varanda, pode codificar quantas mensagens?", Options: "[1] 16 [2] 32 [3] 64 [4] 128 [5] 256: "},
		{Number: "4.", Question: "Bit é uma abreviação de Binary Digit", Options: "[1] Verdadeiro [2] Falso: "},
		{Number: "5.", Question: "On e Off, notação binária, linguagem de máquina, bits, etc., todos referem-se à idéia de que em computadores no fim das contas tudo se resume a zeros e uns", Options: "[1] Verdadeiro [2] Falso: "},
		{Number: "6.", Question: "Mil bytes =", Options: "[1] 1GB [2] 1TB [3] 1KB [4] 1MB [5] isso é pegaginha. Nenhuma das anteriores!: "},
		{Number: "7.", Question: "Mil GB =", Options: "[1] 1KB [2] 1.000.000 bytes [3] 1EB [4] 1GB [5] Isso é pegadinha. Nenhuma das anteriores!: "},
		{Number: "8.", Question: "Quantas 'lâmpadas na varanda' o ENIAC tinha?", Options: "[1] 1600000 [2] 16000 [3] 1600 [4] 16MHz [5] Isso é pegadinha. Nenhuma das anteriores!: "},
		{Number: "9.", Question: "Em Go, rune é um 'apelido' para float64", Options: "[1] Verdadeiro [2] Falso: "},
		{Number: "10.", Question: "Em Go, byte é um 'apelido' para float64", Options: "[1] Verdadeiro [2] Falso: "},
		{Number: "11.", Question: "Ao usar o tipo int, o compilador irá decidir se vai usar int32 ou int64 dependendo da arquitetura do sistema em questão", Options: "[1] Verdadeiro [2] Falso: "},
		{Number: "12.", Question: "Uma string é uma sequência de bytes, também chamada de *fatia de bytes", Options: "[1] Verdadeiro [2] Falso: "},
		{Number: "13.", Question: "O código fonte de Go é sempre UTF-8", Options: "[1] Verdadeiro [2] Falso: "},
		{Number: "14.", Question: "Cada caractere de uma string é um 'rune'", Options: "[1] Verdadeiro [2] Falso: "},
		{Number: "15.", Question: "Via de regra, para tipos numéricos devemos usar float64 para números inteiros, e int para números 'com virgula'", Options: "[1] Verdadeiro [2] Falso: "},
		{Number: "16.", Question: "O número decimal 42 em binário é", Options: "[1] 42 [2] 101010 [3] 2A [4] Ta me batendo uma fome...: "},
		{Number: "17.", Question: "O número decimal 42 em hexadecima é", Options: "[1] 42 [2] 101010 [3] 2A: "},
		{Number: "18.", Question: "Qual das alternativas não representam o número decimal 31337", Options: "[1] 7A69 [2] 7a69 [3] 0x7A69 [4] 0x7a69 [5] Isso é pegadinha, tá de sacanagem?: "},
	}

	gabarito := []format.Answer{
		{Number: "1.", Answer: "1"}, {Number: "2.", Answer: "2"}, {Number: "3.", Answer: "5"}, {Number: "4.", Answer: "1"}, {Number: "5.", Answer: "1"},
		{Number: "6.", Answer: "3"}, {Number: "7.", Answer: "4"}, {Number: "8.", Answer: "2"}, {Number: "9.", Answer: "2"}, {Number: "10.", Answer: "2"},
		{Number: "11.", Answer: "1"}, {Number: "12.", Answer: "1"}, {Number: "13.", Answer: "1"}, {Number: "14.", Answer: "1"}, {Number: "15.", Answer: "2"},
		{Number: "16.", Answer: "2"}, {Number: "17.", Answer: "3"}, {Number: "18.", Answer: "5"},
	}

	sq := new(format.QuestionnaireService)
	sr := new(format.AnswerService)

	prova := sq.Create(questionario)
	respostas, err := sr.Collect(prova)
	if err != nil {
		fmt.Println("Erro ao coletar respostas:", err)
	}

	sr.Validate(respostas, gabarito)
}
