package canais

import "github.com/fabianoflorentino/aprendago/internal/chapter"

func init() {
	chapter.Register(&chapter.Chapter{
		Number:  21,
		Title:   "Canais",
		RootDir: "internal/canais",
	})
}
