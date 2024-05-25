package outline

import "fmt"

func OperadorCurtoDeDeclaracao() {
	operador_curto_de_declaracao := `
- := parece uma marmota (gopher) ou o punisher.
- Uso:
  - Tipagem automática
  - Só pode repetir se houverem variáveis novas
  - != do assignment operator (operador de atribuição)
  - Só funciona dentro de codeblocks
- Terminologia:
  - keywords (palavras-chave) são termos reservados
  - operadores, operandos
  - statement (declaração, afirmação) → uma linha de código, uma instrução que forma uma ação, formada de expressões
  - expressão -> qualquer coisa que "produz um resultado"
  - scope (abrangência)
    - package-level scope
- Lição principal:
  - := utilizado pra criar novas variáveis, dentro de code blocks
  - = para atribuir valores a variáveis já existentes
  `

	fmt.Println("\nOperador de Curto de Declaração")
	fmt.Println(operador_curto_de_declaracao)
}
