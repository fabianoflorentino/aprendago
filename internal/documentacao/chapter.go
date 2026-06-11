package documentacao

import "github.com/fabianoflorentino/aprendago/internal/chapter"

func init() {
	chapter.Register(&chapter.Chapter{
		Number:  25,
		Title:   "Documentação",
		RootDir: "internal/documentacao",
	})
}
