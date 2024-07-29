package funcoes

import "github.com/fabianoflorentino/aprendago/pkg/format"

func Recursividade() {
	topic := format.OutlineContent{
		Title: "Recursividade",
		Content: `
- WP: "The most common application of recursion is in mathematics and computer science, where a function being defined is applied within its own definition."
- Exemplos de recursividade: Fractais, matrioscas, efeito Droste (o efeito produzido por uma imagem que aparece dentro dela própria), GNU (“GNU is Not Unix”), etc.
- No estudo de funções: é uma função que chama a ela própria.
- Exemplo: fatoriais.
  - 4! = 4 * 3 * 2 * 1 (e no zero, deu.)
  - Com recursividade. Go Playground: https://play.golang.org/p/ujsLnUhRp_
  - Com loops. Go Playground: https://play.golang.org/p/F2VsUjYVhc
    `,
	}

	format.FormatOutlineTopic(topic)
}
