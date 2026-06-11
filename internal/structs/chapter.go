package structs

import "github.com/fabianoflorentino/aprendago/internal/chapter"

func init() {
	chapter.Register(&chapter.Chapter{
		Number:  10,
		Title:   "Structs",
		RootDir: "internal/structs",
	})
}
