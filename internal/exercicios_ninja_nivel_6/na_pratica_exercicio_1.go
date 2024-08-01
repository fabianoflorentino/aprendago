package exercicios_ninja_nivel_6

import "github.com/fabianoflorentino/aprendago/pkg/format"

func NaPraticaExercicio1() {
	format.StartSection(func(outline *format.OutlineContent) {
		outline.AddHeader("Na prática: Exercício #1")
		outline.StartList(func(list *format.List) {
			list.AddItem("Exercício:")
			list.StartSubList(func(list *format.List) {
				list.AddItem(("Crie uma função que retorne um int"))
				list.AddItem(("Crie outra função que retorne um int e uma string"))
				list.AddItem(("Chame as duas funções"))
				list.AddItem(("Demonstre seus resultados"))
			})
		})
		outline.StartList(func(list *format.List) {
			list.AddItem("Solução: https://play.golang.org/p/rxJM5fgI-9")
		})

		outline.StartList(func(list *format.List) {
			list.AddItem("Revisão:")
		})
	})
}

func ResolucaoNaPraticaExercicio1() {
}
