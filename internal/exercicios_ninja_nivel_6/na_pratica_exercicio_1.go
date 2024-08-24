package exercicios_ninja_nivel_6

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio1() {
	topic := format.OutlineContent{
		Title: "Na Prática - Exercício 1",
		Content: `
- Exercício:
  - Crie uma função que retorne um int
  - Crie outra função que retorne um int e uma string
  - Chame as duas funções
  - Demonstre seus resultados
- Solução: https://play.golang.org/p/rxJM5fgI-9

- Revisão:
  - Funções!
    - Servem para abstrair código
    - E para reutilizar código
  - A ordem das coisas é:
    - func (receiver) identifier (parameters) (returns) { code }
  - Parâmetros vs. argumentos
  - Funções variádicas
    - Múltiplos parâmetros
    - Múltiplos argumentos
  - Métodos
  - Interfaces & polimorfismo
  - Defer
    - "Deixa pra depois!"
  - Returns
    - Múltiplos returns
    - Returns com nome (blé!)
  - Funcs como expressões
    - Atribuindo uma função a uma variável
  - Callbacks
    - Passando uma função como argumento para outra função
  - Closure
    - Capturando um scope
    - Variáveis declaradas em scopes externos são visíveis em scopes internos
  - Recursividade
    - Uma função que chama a ela mesma
    - Fatoriais
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio1() {
	numero_int := RetornaInt()
	numero, texto := RetornaIntString()

	fmt.Printf("Inteiro: %v\nInteiro&Texto: %v e %v", numero_int, numero, texto)
}

func RetornaInt() int {
	return 42
}

func RetornaIntString() (int, string) {
	return 42, "Olá, mundo!"
}
