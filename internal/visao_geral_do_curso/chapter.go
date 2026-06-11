package visao_geral_do_curso

import "github.com/fabianoflorentino/aprendago/internal/chapter"

func init() {
	chapter.Register(&chapter.Chapter{
		Number:  1,
		Title:   "Visão Geral do Curso",
		RootDir: "internal/visao_geral_do_curso",
	})
}
