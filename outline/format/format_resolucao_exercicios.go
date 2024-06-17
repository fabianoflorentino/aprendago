package format

import "fmt"

func FormatResolucaoExercicios(resolucao any) {
	fmt.Printf("Resolução:\n")
	resolucao = fmt.Sprintf("%v", resolucao)

	fmt.Print(resolucao)
}
