package visao_geral_do_curso

import "github.com/fabianoflorentino/aprendago/pkg/format"

func PorQueGo() {
	format.StartSection(func(outline *format.OutlineContent) {
		outline.AddHeader("Por que Go?")
		outline.StartList(func(list *format.List) {
			list.AddItem("Antes de investir seu tempo em aprender a linguagem Go, é bom você entender por que isso é uma boa idéia.")
			list.AddItem("O que estava acontecendo no Google...")
			list.AddItem("Criada por Ken Thompson (Unix, B, C), Rob Pike (UTF-8), e Robert Griesemer.")
			list.AddItem("Em 2006, não tinha uma linguagem de compilação rápida, execução rápida, e fácil de programar. É uma linguagem criada para resolver as questões de performance e complexidade.")
			list.AddItem("https://golang.org/doc/faq#creating_a_new_language")
			list.AddItem("Eficiente")
			list.StartSubList(func(list *format.List) {
				list.AddItem("Standard library é déis")
				list.AddItem("Multiplataforma.")
				list.AddItem("Garbage collection (lightning fast!)")
				list.AddItem("Cross-compile.")
			})
			list.AddItem("Fácil de usar")
			list.StartSubList(func(list *format.List) {
				list.AddItem("É uma linguagem compilada, de tipagem forte e estática,")
				list.AddItem("Tem pouquíssimas palavras reservadas, que vamos aprender todas no curso, ou seja, é muito de boas de aprender")
				list.AddItem("só sobe nas listas de popularidade")
			})
			list.AddItem("Killer feature: Em 2006, logo após o primeiro dual core. Thread: 1mb. Goroutine: 2kb.")
			list.AddItem("É massa!")
			list.AddItem("Quando usar Go?")
			list.StartSubList(func(list *format.List) {
				list.AddItem("Escala")
				list.AddItem("Seviços web, redes, servers (machine learning, image processing, crypto, ...)")
				list.AddItem("Quando precisar de uma lingaugem rápida, simples, fácil de aprender, e fácil de usar.")
			})
			list.AddItem("Usa em: APIs, CLIs, microservices, libraries/framework, processamento de dados, ... É a base dos serviços de cloud e orquestração de containers.")
			list.AddItem("Quem usa? https://github.com/golang/go/wiki/GoUsers")
			list.AddItem("Go é OOP? https://golang.org/doc/faq#Is_Go_an_object-oriented_language")
			list.AddItem("Mais um: https://golang.org/doc/faq#principles")
			list.AddItem("$$$: https://insights.stackoverflow.com/survey/2017#technology-top-paying-technologies-by-region -> US")
		})
	})
}
