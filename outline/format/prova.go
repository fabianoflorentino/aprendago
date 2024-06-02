package outline

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
}

func CriarQuestionario(questionario []Questionario) []Questionario {
	var prova []Questionario

	for _, quest := range questionario {
		prova = append(prova, Questionario{Numero: quest.Numero, Pergunta: quest.Pergunta, Opcoes: quest.Opcoes})
	}

	return prova
}

func ColetarRespostas(questionario []Questionario) ([]Resposta, error) {
	var listaRespostas []Resposta
	var resposta string

	criar_questionario := CriarQuestionario(questionario)

	for _, p := range criar_questionario {
		fmt.Println(p.Numero, p.Pergunta, p.Opcoes)
		fmt.Println(p.Opcoes)
		fmt.Scan(&resposta)

		if !opcaoValida(resposta, p.Opcoes) {
			return nil, errors.New("opção inválida")
		}

		listaRespostas = append(listaRespostas, Resposta{p.Numero, resposta})

	}

	return listaRespostas, nil
}

func opcaoValida(resposta, opcoes string) bool {
	for _, opt := range strings.Fields(opcoes) {
		opt = strings.Trim(opt, "[]")
		if resposta == opt {
			return true
		}
	}

	return false
}

func ImprimirRespostas(respostas []Resposta) {
	fmt.Printf("\nRespostas da prova:\n\n")

	numRespostas := len(respostas)
	for i := 0; i <= numRespostas; i++ {
		if i+1 <= numRespostas {
			fmt.Printf("P: %-4s R: %-8s\n", respostas[i].Numero, respostas[i].Resposta)
		}
	}
}

// ImprimiGabarito exibe o gabarito da prova
func ImprimiGabarito(gabarito []Resposta) {
	// Exibe o gabarito da prova para o usuário em  duas colunas de 15 perguntas com suas respectivas respostas
	fmt.Printf("\nGabarito da prova:\n\n")
	for i := 0; i <= len(gabarito); i++ {
		if i+1 <= len(gabarito) {
			fmt.Printf("P: %-4s R: %-8s\n", gabarito[i].Numero, gabarito[i].Resposta)
		}
	}
}
