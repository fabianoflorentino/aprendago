package visao_geral_do_curso

import "github.com/fabianoflorentino/aprendago/pkg/format"

func Recursos() {
	format.StartSection(func(outline *format.OutlineContent) {
		outline.AddHeader("Recursos")
		outline.StartList(func(list *format.List) {
			list.AddItem("Leia as descrições dos vídeos!")
			list.StartSubList(func(list *format.List) {
				list.AddItem("Outline completo: https://github.com/ellenkorbes/aprendago/blob/master/OUTLINE.md")
			})
			list.AddItem("Comunidade:")
			list.StartSubList(func(list *format.List) {
				list.AddItem("Slack: https://invite.slack.golangbridge.org/")
				list.AddItem("#brazil")
				list.AddItem("#brazilian-go-studies, toda quinta às 22h!")
				list.StartSubList(func(list *format.List) {
					list.AddItem("gravações em: https://www.youtube.com/cesargimenes")
					list.AddItem("exercícios em: https://github.com/crgimenes/Go-Hands-On")
					list.AddItem("FB: https://www.facebook.com/gophers.br/")
				})
			})
			list.AddItem("Livros:")
			list.StartSubList(func(list *format.List) {
				list.AddItem("A Linguagem de Programação Go (Alan Donovan, Brian Kernighan): https://www.amazon.com.br/dp/8575225464/")
				list.AddItem("Go em Ação (William Kennedy, Brian Ketelsen, Erik St. Martin): https://www.amazon.com.br/dp/8575225065/")
				list.AddItem("Introdução à Linguagem Go (Caleb Doxsey): https://www.amazon.com.br/dp/8575224891/")
			})
			list.AddItem("Em inglês: gobyexample.com, forum.golangbridge.org, stackoverflow.com/questions/tagged/go")
			list.AddItem("Documentação oficial:")
			list.StartSubList(func(list *format.List) {
				list.AddItem("golang.org")
				list.AddItem("godoc.org")
				list.AddItem("golang.org/doc/effective_go.html")
			})
			list.AddItem("Podcast Go Time: https://changelog.com/gotime")
			list.AddItem("Defective Go: https://www.youtube.com/channel/UC98qIvCCqd4fjOw1ks78SwA.")
		})
	})
}
