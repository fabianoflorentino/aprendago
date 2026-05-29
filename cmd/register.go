package cmd

import (
	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

// init registers migrated chapters with the chapter registry.
//
// During migration (Phases 1-3), chapters are registered here because the old
// internal/chapter/chapter.go imports all chapter packages, which would create
// an import cycle if chapters imported internal/chapter in their own init().
//
// After Fase 4 cleanup (when old internal/chapter/chapter.go is removed),
// each chapter will register itself via its own init() in
// internal/<chapter>/chapter.go.
func init() {
	chapter.Register(&chapter.Chapter{
		Number:  1,
		Title:   "Visão Geral do Curso",
		RootDir: "internal/visao_geral_do_curso",
	})
}
