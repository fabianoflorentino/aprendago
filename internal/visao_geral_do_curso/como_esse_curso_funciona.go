package visao_geral_do_curso

import "github.com/fabianoflorentino/aprendago/pkg/format"

func ComoEsseCursoFunciona() {
	format.StartSection(func(outline *format.OutlineContent) {
		outline.AddHeader("Como esse curso funciona")
		outline.StartList(func(list *format.List) {
			list.AddItem("Velocidade de playback")
			list.AddItem("Repetição.")
			list.AddItem("Erros.")
			list.AddItem("Português vs. inglês")
			list.StartSubList(func(list *format.List) {
				list.AddItem("Obviamente esse é um curso em português.")
				list.AddItem("Mas a língua semi-oficial da programação e do mundo da tecnologia em geral é a língua inglesa.")
				list.AddItem("Por isso vamos utilizar bastante termos em inglês ao longo do curso.")
				list.AddItem("As explicações serão em português, claro, mas quero que todos fiquem bem confortáveis com os termos em inglês.")
				list.AddItem("Desta maneira você será auto-suficiente e poderá procurar por documentação e ajuda por toda a internet, não somente nos sites em português.")
				list.AddItem("(Alem de que eu não sei e nem quero saber o nome em português de boa parte dessas coisas :P)")
			})
			list.AddItem(`Constantemente "em progresso."`)
			list.AddItem("Seu feedback é super importante!")
		})
	})
}
