package outline

import (
	"fmt"

	helpme "github.com/fabianoflorentino/aprendago/outline/format"
)

func HelpMeVisaoGeralDoCurso() {
	hlp := []helpme.HelpMe{
		{Flag: "--bem-vindo", Description: "Exibe a mensagem de boas-vindas ao curso Aprenda Go.", Width: 0},
		{Flag: "--porque-go", Description: "Descreve os benefícios e razões para aprender a linguagem Go.", Width: 0},
		{Flag: "--sucesso", Description: "Apresenta dicas e estratégias para ter sucesso no curso.", Width: 0},
		{Flag: "--recursos", Description: "Lista recursos e materiais de apoio para o curso.", Width: 0},
		{Flag: "--como-esse-curso-funciona", Description: "Explica a estrutura e metodologia do curso.", Width: 0},
	}

	fmt.Println("Capítulo 1: Visão Geral do Curso")
	helpme.PrintHelpMe(hlp)
}
