package format

import "fmt"

func FormatResolucaoExercicios(resolucao any) {
	fmt.Print("Resolução:\n")
	resolucao = fmt.Sprintf("%v", resolucao)

	fmt.Print(resolucao)
}
