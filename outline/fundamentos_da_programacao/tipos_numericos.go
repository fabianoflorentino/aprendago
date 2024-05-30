package outline

import "fmt"

func TiposNumericos() {
	tipos_numericos := `
- int vs. float: Números inteiros vs. números com frações.
- golang.org/ref/spec → numeric types
- Integers:
  - Números inteiros
  - int & uint → “implementation-specific sizes”
  - Todos os tipos numéricos são distintos, exceto:
    - byte = uint8
    - rune = int32 (UTF8)
    (O código fonte da linguagem Go é sempre em UTF-8).
  - Tipos são únicos
    - Go é uma linguagem estática
    - int e int32 não são a mesma coisa
    - Para "misturá-los" é necessário conversão
  - Regra geral: use somente int
- Floating point:
  - Números racionais ou reais
  - Regra geral: use somente float64
- Na prática:
  - Defaults com :=
  - Tipagem com var
  - Dá pra colocar número com vírgula em tipo int?
  - Overflow
  - Go Playground: https://play.golang.org/p/dt2x1ies5b
- “implementation-specific sizes”? Runtime package. Word.
  - GOOS
  - GORUNTIME
  - https://play.golang.org/p/1vp5DImIMM
  `

	fmt.Println("Tipos numéricos")
	fmt.Println(tipos_numericos)
}
