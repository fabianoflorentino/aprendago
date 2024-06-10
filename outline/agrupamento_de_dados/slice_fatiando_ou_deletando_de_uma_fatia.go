package agrupamento_de_dados

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/outline/format"
)

func SliceFatiandoOuDeletandoDeUmaFatia() {
	topic := format.OutlineContent{
		Title: "Slice: fatiando ou deletando de uma fatia",
		Content: `
- x[:], x[a:], x[:b], x[a:b]
- "a" é incluso;
- "b" não é.
- Exemplo: cabeça magnética de um disco rígido (relógio, fita).
  - Off-by-one error.
- Go Playground: https://play.golang.org/p/i5ZOLKb3Fi
- É fatiando que se deleta um item de uma slice. Na prática:
  - x := append(x[:i], x[:i]...)
  - Go Playground: https://play.golang.org/p/xK2HwCqvwd
- Exercício: tente acessar todos os itens de uma slice *sem* utilizar range.
- Solução: https://play.golang.org/p/aUC9qVCobH ou use no menu do programa (--slice-fatiando-ou-deletando-de-uma-fatia --resolucao)
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoFatiaOuDeletandoDeUmaFatia() {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(slice[:]) // [1 2 3 4 5]
}
