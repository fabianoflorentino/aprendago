package exercicios_ninja_nivel_2

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func ResolucaoNaPraticaExercicio1() {
	numero := 1000019

	resoluca := fmt.Sprintf("Decimal: %d \nBinário: %b \nHexadecimal: %#x \n", numero, numero, numero)

	format.FormatResolucaoExercicios(resoluca)
}

func ResolucaoNaPraticaExercicio2() {
	a := 10
	b := 20

	resolucao := fmt.Sprintf("a: %v \nb: %v \n\na == b: %v \na != b: %v \na <= b: %v \na < b: %v \na >= b: %v \na > b: %v \n", a, b, a == b, a != b, a <= b, a < b, a >= b, a > b)

	format.FormatResolucaoExercicios(resolucao)
}

func ResolucaoNaPraticaExercicio3() {
	const (
		untypedConst     = 10
		typedConst   int = 20
	)

	resolucao := fmt.Sprintf("untypedConst: %v \ntypedConst: %v", untypedConst, typedConst)

	format.FormatResolucaoExercicios(resolucao)
}

func ResolucaoNaPraticaExercicio4() {
	v1 := 10

	resolucao1 := fmt.Sprintf("Decimal: %d \nBinário: %b \nHexadecimal: %#x \n\n", v1, v1, v1)

	v2 := v1 << 1

	resolucao2 := fmt.Sprintf("Decimal: %d \nBinário: %b \nHexadecimal: %#x", v2, v2, v2)

	format.FormatResolucaoExercicios(resolucao1 + resolucao2)
}

func ResolucaoNaPraticaExercicio5() {
	v1 := `Aprenda Go!`

	format.FormatResolucaoExercicios(v1)
}

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
