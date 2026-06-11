package agrupamento_de_dados

import "github.com/fabianoflorentino/aprendago/internal/chapter"

func init() {
	chapter.Register(&chapter.Chapter{
		Number:  8,
		Title:   "Agrupamento de Dados",
		RootDir: "internal/agrupamento_de_dados",
	})
}
