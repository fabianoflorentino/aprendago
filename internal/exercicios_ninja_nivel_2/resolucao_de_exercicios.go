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
	questionario := []format.Questionario{
		{Numero: "1.", Pergunta: "Ao utilizar o opeardor curto de declaração, não é necessário especificar o tipo", Opcoes: "[1] Verdadeiro [2] Falso: "},
		{Numero: "2.", Pergunta: "Uma variável do tipo bool pode conter três valores: true, false, undefined", Opcoes: "[1] Verdadeiro [2] Falso: "},
		{Numero: "3.", Pergunta: "Se você tiver 5 lâmpadas na varanda, pode codificar quantas mensagens?", Opcoes: "[1] 16 [2] 32 [3] 64 [4] 128 [5] 256: "},
		{Numero: "4.", Pergunta: "Bit é uma abreviação de Binary Digit", Opcoes: "[1] Verdadeiro [2] Falso: "},
		{Numero: "5.", Pergunta: "On e Off, notação binária, linguagem de máquina, bits, etc., todos referem-se à idéia de que em computadores no fim das contas tudo se resume a zeros e uns", Opcoes: "[1] Verdadeiro [2] Falso: "},
		{Numero: "6.", Pergunta: "Mil bytes =", Opcoes: "[1] 1GB [2] 1TB [3] 1KB [4] 1MB [5] isso é pegaginha. Nenhuma das anteriores!: "},
		{Numero: "7.", Pergunta: "Mil GB =", Opcoes: "[1] 1KB [2] 1.000.000 bytes [3] 1EB [4] 1GB [5] Isso é pegadinha. Nenhuma das anteriores!: "},
		{Numero: "8.", Pergunta: "Quantas 'lâmpadas na varanda' o ENIAC tinha?", Opcoes: "[1] 1600000 [2] 16000 [3] 1600 [4] 16MHz [5] Isso é pegadinha. Nenhuma das anteriores!: "},
		{Numero: "9.", Pergunta: "Em Go, rune é um 'apelido' para float64", Opcoes: "[1] Verdadeiro [2] Falso: "},
		{Numero: "10.", Pergunta: "Em Go, byte é um 'apelido' para float64", Opcoes: "[1] Verdadeiro [2] Falso: "},
		{Numero: "11.", Pergunta: "Ao usar o tipo int, o compilador irá decidir se vai usar int32 ou int64 dependendo da arquitetura do sistema em questão", Opcoes: "[1] Verdadeiro [2] Falso: "},
		{Numero: "12.", Pergunta: "Uma string é uma sequência de bytes, também chamada de *fatia de bytes", Opcoes: "[1] Verdadeiro [2] Falso: "},
		{Numero: "13.", Pergunta: "O código fonte de Go é sempre UTF-8", Opcoes: "[1] Verdadeiro [2] Falso: "},
		{Numero: "14.", Pergunta: "Cada caractere de uma string é um 'rune'", Opcoes: "[1] Verdadeiro [2] Falso: "},
		{Numero: "15.", Pergunta: "Via de regra, para tipos numéricos devemos usar float64 para números inteiros, e int para números 'com virgula'", Opcoes: "[1] Verdadeiro [2] Falso: "},
		{Numero: "16.", Pergunta: "O número decimal 42 em binário é", Opcoes: "[1] 42 [2] 101010 [3] 2A [4] Ta me batendo uma fome...: "},
		{Numero: "17.", Pergunta: "O número decimal 42 em hexadecima é", Opcoes: "[1] 42 [2] 101010 [3] 2A: "},
		{Numero: "18.", Pergunta: "Qual das alternativas não representam o número decimal 31337", Opcoes: "[1] 7A69 [2] 7a69 [3] 0x7A69 [4] 0x7a69 [5] Isso é pegadinha, tá de sacanagem?: "},
	}

	gabarito := []format.Resposta{
		{Numero: "1.", Resposta: "1"}, {Numero: "2.", Resposta: "2"}, {Numero: "3.", Resposta: "5"}, {Numero: "4.", Resposta: "1"}, {Numero: "5.", Resposta: "1"},
		{Numero: "6.", Resposta: "3"}, {Numero: "7.", Resposta: "4"}, {Numero: "8.", Resposta: "2"}, {Numero: "9.", Resposta: "2"}, {Numero: "10.", Resposta: "2"},
		{Numero: "11.", Resposta: "1"}, {Numero: "12.", Resposta: "1"}, {Numero: "13.", Resposta: "1"}, {Numero: "14.", Resposta: "1"}, {Numero: "15.", Resposta: "2"},
		{Numero: "16.", Resposta: "2"}, {Numero: "17.", Resposta: "3"}, {Numero: "18.", Resposta: "5"},
	}

	sq := new(format.ServicoQuestionario)
	sr := new(format.ServicoResposta)

	prova := sq.Cria(questionario)
	respostas, err := sr.Coleta(prova)
	if err != nil {
		fmt.Println("Erro ao coletar respostas:", err)
	}

	sr.Valida(respostas, gabarito)
}
