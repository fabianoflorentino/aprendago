package seu_ambiente_de_desenvolvimento

import "github.com/fabianoflorentino/aprendago/internal/chapter"

func init() {
	chapter.Register(&chapter.Chapter{
		Number:  19,
		Title:   "Seu Ambiente de Desenvolvimento",
		RootDir: "internal/seu_ambiente_de_desenvolvimento",
	})
}
