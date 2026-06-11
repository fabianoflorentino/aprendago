package fundamentos_da_programacao

import "github.com/fabianoflorentino/aprendago/internal/chapter"

func init() {
	chapter.Register(&chapter.Chapter{
		Number:  4,
		Title:   "Fundamentos da Programação",
		RootDir: "internal/fundamentos_da_programacao",
	})
}
