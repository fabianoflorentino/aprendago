package outline

import "fmt"

func ValorZero() {
	valor_zero := `
- Declaração vs. inicialização vs. atribuição de valor. Variáveis: caixas postais.
- O que é valor zero?
- Os zeros:
  - ints: 0
  - floats: 0.0
  - booleans: false
  - strings: ""
  - pointers, functions, interfaces, slices, channels, maps: nil
- Use := sempre que possível.
- Use var para package-level scope.
  `

	fmt.Println("\nValor Zero")
	fmt.Println(valor_zero)
}
