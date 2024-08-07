package variaveis_valores_tipos

import "github.com/fabianoflorentino/aprendago/pkg/format"

func OperadorCurtoDeDeclaracao() {
	topic := format.OutlineContent{
		Title: "Operador curto de declaração",
		Content: `
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
    `,
	}

	format.FormatOutlineTopic(topic)
}
