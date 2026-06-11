package aplicacoes

import "github.com/fabianoflorentino/aprendago/internal/chapter"

func init() {
	chapter.Register(&chapter.Chapter{
		Number:  16,
		Title:   "Aplicações",
		RootDir: "internal/aplicacoes",
	})
}
