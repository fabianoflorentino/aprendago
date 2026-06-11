package teste_benchmarking

import "github.com/fabianoflorentino/aprendago/internal/chapter"

func init() {
	chapter.Register(&chapter.Chapter{
		Number:  27,
		Title:   "Teste e Benchmarking",
		RootDir: "internal/teste_benchmarking",
	})
}
