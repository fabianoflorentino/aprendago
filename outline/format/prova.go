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
	Status   string
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
		fmt.Print(p.Numero, p.Pergunta, p.Opcoes)
		fmt.Scan(&resposta)

		if !opcaoValida(resposta, p.Opcoes) {
			return nil, errors.New("opção inválida")
		}

		listaRespostas = append(listaRespostas, Resposta{p.Numero, resposta, ""})
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

func ValidarRespostas(respostas []Resposta, gabarito []Resposta) {
	var validaRespostas []Resposta
	var totalCorretas int

	for i, resp := range respostas {
		if resp.Resposta == gabarito[i].Resposta {
			validaRespostas = append(validaRespostas, Resposta{resp.Numero, resp.Resposta, "Correta"})
			totalCorretas++
		} else {
			validaRespostas = append(validaRespostas, Resposta{resp.Numero, resp.Resposta, "Incorreta"})
		}
	}

	fmt.Printf("\nResultado da prova:\n")
	fmt.Printf("\nTotal de respostas corretas: %d de %d\n", totalCorretas, len(gabarito))

	resultadoDaProva := float64(totalCorretas) / float64(len(gabarito)) * 100
	fmt.Printf("Percentual de acerto: %.2f%%.\n\n", resultadoDaProva)

	for _, v := range validaRespostas {
		fmt.Printf("P: %-4s R: %-4s %s\n", v.Numero, v.Resposta, v.Status)
	}
}
