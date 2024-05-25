package outline

import "fmt"

func ApalavraReservadaVar() {
	a_palava_reservada_var := `
- Variável declarada em um code block é undefined em outro
- Para variáveis com uma abrangência maior, package level scope, utilizamos "var"
- Funciona em qualquer lugar
- Prestar atenção: chaves, colchetes, parênteses
  `

	fmt.Println("\nA palavra reservada var")
	fmt.Println(a_palava_reservada_var)
}
