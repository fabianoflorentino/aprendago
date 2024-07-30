package visao_geral_do_curso

import "github.com/fabianoflorentino/aprendago/pkg/format"

func BemVindo() {
	format.StartSection(func(outline *format.OutlineContent) {
		outline.AddHeader("Bem-vindo")
		outline.StartList(func(list *format.List) {
			list.AddItem("Bem vindo ao curso!")
			list.AddItem("Eu sou...")
			list.AddItem("Go foi criado por gente foda que criou o Unix, B, UTF-8...")
			list.AddItem("Em 2006 o google queria...")
			list.AddItem("É uma lingguagem que vem crescrendo horrores...")
			list.AddItem("Nesse curso você vai aprender...")
			list.AddItem("O curriculo que vamos estudar...")
			list.AddItem("Para os novos na programação... Para os programadores experientes...")
			list.AddItem("Participe!")
		})
	})
}
