package format

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

// Template para formatar a saída de ajuda. O modelo usa a função printf para formatar a saída.
var templateHelpMe = `
{{- range .}}
  {{ printf "%-*s" .Width .Flag }}   {{ .Description | indent .Width}}
{{- end }}
`

// HelpMe é uma estrutura que representa uma opção de ajuda.
type HelpMe struct {
	Flag        string
	Description string
	Width       int
}

// parseWidth é uma função que calcula o tamanho máximo de uma flag. O tamanho máximo é usado para formatar a saída de ajuda.
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
func PrintHelpMe(helpme []HelpMe) {
	width := parseWidth(helpme)
	for i := range helpme {
		helpme[i].Width = width
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
	err = tmpl.Execute(os.Stdout, helpme)
	if err != nil {
		panic(err)
	}
}
