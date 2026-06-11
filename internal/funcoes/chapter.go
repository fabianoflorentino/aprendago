package funcoes

import "github.com/fabianoflorentino/aprendago/internal/chapter"

func init() {
	chapter.Register(&chapter.Chapter{
		Number:  12,
		Title:   "Funções",
		RootDir: "internal/funcoes",
	})
}
