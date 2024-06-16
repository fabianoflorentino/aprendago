package exercicios_ninja_nivel_2

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio7() {
	topic := format.OutlineContent{
		Title: "Na Prática - Exercício #7",
		Content: `
- Prova!
- Link: https://goo.gl/forms/fnPXMmxvKAEUD8xP2 [Link quebrado]. (Use a opção --na-pratica-exercicio-7 --nivel-2 --prova)
- Motivação. Ninja nível 2!
    `,
	}

	format.FormatOutlineTopic(topic)
}

func RespondaAProva() {
	questionario := []format.Questionario{
		{Numero: "1.", Pergunta: "Ao utilizar o opeardor curto de declaração, não é necessário especificar o tipo", Opcoes: []string{"Verdade", "Mentira"}},
		{Numero: "2.", Pergunta: "Uma variável do tipo bool pode conter três valores: true, false, undefined", Opcoes: []string{"Verdade", "Mentira"}},
		{Numero: "3.", Pergunta: "Se você tiver 5 lâmpadas na varanda, pode codificar quantas mensagens?", Opcoes: []string{"16", "32", "64", "128", "256"}},
		{Numero: "4.", Pergunta: "Bit é uma abreviação de Binary Digit", Opcoes: []string{"Verdade", "Mentira"}},
		{Numero: "5.", Pergunta: "On e Off, notação binária, linguagem de máquina, bits, etc., todos referem-se à idéia de que em computadores no fim das contas tudo se resume a zeros e uns", Opcoes: []string{"Verdade", "Mentira"}},
		{Numero: "6.", Pergunta: "Mil bytes =", Opcoes: []string{"1GB", "1TB", "1KB", "1MB", "isso é pegaginha. Nenhuma das anteriores!"}},
		{Numero: "7.", Pergunta: "Mil GB =", Opcoes: []string{"1KB", "1.000.000 bytes", "1EB", "1GB", "Isso é pegadinha. Nenhuma das anteriores!"}},
		{Numero: "8.", Pergunta: "Quantas 'lâmpadas na varanda' o ENIAC tinha?", Opcoes: []string{"1600000", "16000", "1600", "16MHz", "Isso é pegadinha. Nenhuma das anteriores!"}},
		{Numero: "9.", Pergunta: "Em Go, rune é um 'apelido' para float64", Opcoes: []string{"Verdade", "Mentira"}},
		{Numero: "10.", Pergunta: "Em Go, byte é um 'apelido' para float64", Opcoes: []string{"Verdade", "Mentira"}},
		{Numero: "11.", Pergunta: "Ao usar o tipo int, o compilador irá decidir se vai usar int32 ou int64 dependendo da arquitetura do sistema em questão", Opcoes: []string{"Verdade", "Mentira"}},
		{Numero: "12.", Pergunta: "Uma string é uma sequência de bytes, também chamada de *fatia de bytes", Opcoes: []string{"Verdade", "Mentira"}},
		{Numero: "13.", Pergunta: "O código fonte de Go é sempre UTF-8", Opcoes: []string{"Verdade", "Mentira"}},
		{Numero: "14.", Pergunta: "Cada caractere de uma string é um 'rune'", Opcoes: []string{"Verdade", "Mentira"}},
		{Numero: "15.", Pergunta: "Via de regra, para tipos numéricos devemos usar float64 para números inteiros, e int para números 'com virgula'", Opcoes: []string{"Verdade", "Mentira"}},
		{Numero: "16.", Pergunta: "O número decimal 42 em binário é", Opcoes: []string{"42", "101010", "2A", "Ta me batendo uma fome..."}},
		{Numero: "17.", Pergunta: "O número decimal 42 em hexadecima é", Opcoes: []string{"42", "101010", "2A"}},
		{Numero: "18.", Pergunta: "Qual das alternativas não representam o número decimal 31337", Opcoes: []string{"7A69", "7a69", "0x7A69", "0x7a69", "Isso é pegadinha, tá de sacanagem?"}},
	}

	gabarito := []format.Resposta{
		{Numero: "1.", Resposta: 1}, {Numero: "2.", Resposta: 2}, {Numero: "3.", Resposta: 5}, {Numero: "4.", Resposta: 1}, {Numero: "5.", Resposta: 1},
		{Numero: "6.", Resposta: 3}, {Numero: "7.", Resposta: 4}, {Numero: "8.", Resposta: 2}, {Numero: "9.", Resposta: 2}, {Numero: "10.", Resposta: 2},
		{Numero: "11.", Resposta: 1}, {Numero: "12.", Resposta: 1}, {Numero: "13.", Resposta: 1}, {Numero: "14.", Resposta: 1}, {Numero: "15.", Resposta: 2},
		{Numero: "16.", Resposta: 2}, {Numero: "17.", Resposta: 3}, {Numero: "18.", Resposta: 5},
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
