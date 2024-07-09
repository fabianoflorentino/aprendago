package agrupamento_de_dados

import "github.com/fabianoflorentino/aprendago/pkg/format"

func SliceMake() {
	topic := format.OutlineContent{
		Title: "Slice: Make",
		Content: `
- Slices são feitas de arrays.
- Elas são dinâmicas, podem mudar de tamanho.
- Sempre que isso acontece, um novo array é criado e os dados são copiados.
- É conveniente, mas tem um custo computacional.
- Para otimizar as coisas, podemos utilizar make.
- make([]T, len, cap)
- "The length of a slice may be changed as long as it still fits within the limits of the underlying array; just assign it to a slice of itself. The capacity of a slice, accessible by the built-in function cap, reports the maximum length the slice may assume."
- len(x), cap(x)
- x[n] onde n > len é out of range. Use append.
- Append > cap modifica o array subjacente.
- pkg/builtin/#append: "If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated."
- Effective Go.
- Go Playground: https://play.golang.org/p/e8GWzyEEL8
    `,
	}

	format.FormatOutlineTopic(topic)
}
