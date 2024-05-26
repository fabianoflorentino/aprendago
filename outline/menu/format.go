package outline

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// {{ printf "%-*s" .Width .Flag }}: Aqui, estamos usando a sintaxe de template do Go, que permite a inserção de valores
// dinâmicos em uma string formatada. A função printf é usada para formatar a string. O formato da string é "%-*s", que
// é uma diretiva de formatação que especifica que o valor a ser inserido será uma string (%s).
// O sinal de menos (-) indica que a string será alinhada à esquerda. O valor .* é uma notação especial que indica que o
// valor será fornecido como argumento adicional. Nesse caso, os argumentos adicionais são .Width e .Flag
// O código usa a função printf para formatar uma string, inserindo valores dinâmicos como .Width, .Flag e .Description.
// O resultado formatado da string será impresso ou usado de alguma outra forma no restante do código.
var templateHelpMe = `
Uso: go run main.go [opção]

Exemplos:
  go run main.go --bem-vindo

Opções:

{{- range .}}
  {{ printf "%-*s" .Width .Flag }}   {{ .Description | indent .Width}}
{{- end }}
`

// HelpMe é uma estrutura que representa uma opção disponível para o programa.
// Cada opção tem uma flag, uma descrição e um tamanho.
type HelpMe struct {
	Flag        string
	Description string
	Width       int
}

// HELPME é uma lista de opções disponíveis para o programa.
// Cada opção é representada por uma estrutura HelpMe que contém a flag, a descrição e o tamanho da flag.
var HELPME = []HelpMe{
	{"--bem-vindo", "Exibe a mensagem de boas-vindas ao curso Aprenda Go.", 0},
	{"--porque-go", "Descreve os benefícios e razões para aprender a linguagem Go.", 0},
	{"--sucesso", "Apresenta dicas e estratégias para ter sucesso no curso.", 0},
	{"--recursos", "Lista recursos e materiais de apoio para o curso.", 0},
	{"--como-esse-curso-funciona", "Explica a estrutura e metodologia do curso.", 0},
	{"--go-playground", "Fornece informações sobre o uso do Go Playground.", 0},
	{"--hello-world", "Detalha o primeiro programa em Go: 'Hello, World!'.", 0},
	{"--operador-curto-de-declaracao", "Explica o uso do operador curto de declaração em Go.", 0},
	{"--a-palavra-reservada-var", "Descreve o uso da palavra reservada 'var' em Go.", 0},
	{"--explorando-tipos", "Explora os diferentes tipos de dados em Go.", 0},
	{"--valor-zero", "Discute o conceito de valor zero em Go.", 0},
	{"--o-pacote-fmt", "Fornece um resumo do pacote 'fmt' para formatação de E/S.", 0},
	{"--criando-seu-proprio-tipo", "Detalha como criar seus próprios tipos em Go.", 0},
	{"--conversao-nao-coercao", "Explica a diferença entre conversão e coerção em Go.", 0},
	{"--contribua-seu-codigo", "Fornece informações sobre como contribuir com seu próprio código.", 0},
	{"--outline", "Exibe o outline completo do curso.", 0},
	{"--help", "Exibe a lista de todas as opções disponíveis.", 0},
}

// parseWidth é uma função que recebe uma lista de flags e retorna o tamanho da maior flag
// para que possamos formatar a saída de ajuda de maneira uniforme.
func parseWidth(flags []HelpMe) int {
	maxWidth := 0

	// O loop for percorre a lista de flags e verifica o tamanho de cada flag.
	// Se o tamanho da flag for maior que o tamanho máximo, o tamanho máximo é atualizado.
	for _, flag := range flags {
		flag.Flag = strings.TrimSpace(flag.Flag)
		if len(flag.Flag) > maxWidth {
			maxWidth = len(flag.Flag)
		}
	}

	return maxWidth
}

// indent adiciona um recuo à frente de cada linha de uma string.
// O recuo é calculado com base no tamanho da flag e em um valor fixo (4).
func indent(width int, description string) string {
	lines := strings.Split(strings.TrimSpace(description), "\n")
	if len(lines) <= 1 {
		return description
	}

	// O loop for percorre as linhas da descrição e adiciona um recuo à frente de cada linha.
	// O recuo é calculado com base no tamanho da flag e em um valor fixo (4).
	for idx := 1; idx < len(lines); idx++ {
		lines[idx] = strings.Repeat(" ", width+4) + lines[idx]
	}

	return strings.Join(lines, "\n")
}

// PrintHelpMe é uma função que imprime a lista de opções disponíveis para o programa.
// A função usa um template para formatar a saída de ajuda de maneira uniforme.
func PrintHelpMe() {
	width := parseWidth(HELPME)
	for i := range HELPME {
		HELPME[i].Width = width
	}

	// funcMap é um mapa de funções que podem ser usadas no template.
	// Neste caso, estamos usando a função printf do pacote fmt para formatar a saída.
	funcMap := template.FuncMap{
		"printf": fmt.Sprintf,
		"indent": indent,
	}

	// tmpl é um objeto de template que contém o modelo de saída de ajuda.
	// O modelo é uma string que contém a formatação da saída de ajuda.
	tmpl, err := template.New("helpme").Funcs(funcMap).Parse(templateHelpMe)
	if err != nil {
		panic(err)
	}

	// A função Execute é usada para renderizar o modelo e imprimir a saída formatada.
	// O primeiro argumento é o destino da saída (os.Stdout) e o segundo argumento é o modelo de dados (HELPME).
	err = tmpl.Execute(os.Stdout, HELPME)
	if err != nil {
		panic(err)
	}
}
