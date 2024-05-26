package outline

import (
	"fmt"
)

func NaPraticaExercicio2() {
	na_pratica_exercicio_2 := `
- Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:
  1. Identificador "x" deverá ter tipo int
  2. Identificador "y" deverá ter tipo string
  3. Identificador "z" deverá ter tipo bool
- Na função main:
  1. Demonstre os valores de cada identificador
  2. O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?
- Solução: https://play.golang.org/p/pAFd-F7uGZ
  `

	fmt.Println("Na prática: Exercicio #2")
	fmt.Println(na_pratica_exercicio_2)
}

func ResolucaoNaPraticaExercicio2() {
	var x int
	var y string
	var z bool

	fmt.Println("Valor de x:", x)
	fmt.Println("Valor de y:", y)
	fmt.Println("Valor de z:", z)
}
