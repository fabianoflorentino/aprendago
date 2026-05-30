package cmd

import (
	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

// init registers all chapters with the chapter registry.
//
// During migration (Phases 1-3), chapters are registered here because the old
// internal/chapter/chapter.go imports all chapter packages, which would create
// an import cycle if chapters imported internal/chapter in their own init().
//
// After Fase 4 cleanup (when old internal/chapter/chapter.go is removed),
// each chapter will register itself via its own init() in
// internal/<chapter>/chapter.go.
func init() {
	chapters := []struct {
		number  int
		title   string
		rootDir string
	}{
		{1, "Visão Geral do Curso", "internal/visao_geral_do_curso"},
		{2, "Variáveis, Valores e Tipos", "internal/variaveis_valores_tipos"},
		{3, "Exercícios Ninja: Nível 1", "internal/exercicios_ninja_nivel_1"},
		{4, "Fundamentos da Programação", "internal/fundamentos_da_programacao"},
		{5, "Exercícios Ninja: Nível 2", "internal/exercicios_ninja_nivel_2"},
		{6, "Fluxo de Controle", "internal/fluxo_de_controle"},
		{7, "Exercícios Ninja: Nível 3", "internal/exercicios_ninja_nivel_3"},
		{8, "Agrupamento de Dados", "internal/agrupamento_de_dados"},
		{9, "Exercícios Ninja: Nível 4", "internal/exercicios_ninja_nivel_4"},
		{10, "Structs", "internal/structs"},
		{11, "Exercícios Ninja: Nível 5", "internal/exercicios_ninja_nivel_5"},
		{12, "Funções", "internal/funcoes"},
		{13, "Exercícios Ninja: Nível 6", "internal/exercicios_ninja_nivel_6"},
		{14, "Ponteiros", "internal/ponteiros"},
		{15, "Exercícios Ninja: Nível 7", "internal/exercicios_ninja_nivel_7"},
		{16, "Aplicações", "internal/aplicacoes"},
		{17, "Exercícios Ninja: Nível 8", "internal/exercicios_ninja_nivel_8"},
		{18, "Concorrência", "internal/concorrencia"},
		{19, "Seu Ambiente de Desenvolvimento", "internal/seu_ambiente_de_desenvolvimento"},
		{20, "Exercícios Ninja: Nível 9", "internal/exercicios_ninja_nivel_9"},
		{21, "Canais", "internal/canais"},
		{22, "Exercícios Ninja: Nível 10", "internal/exercicios_ninja_nivel_10"},
		{23, "Tratamento de Erro", "internal/tratamento_de_erro"},
		{24, "Exercícios Ninja: Nível 11", "internal/exercicios_ninja_nivel_11"},
		{25, "Documentação", "internal/documentacao"},
		{26, "Exercícios Ninja: Nível 12", "internal/exercicios_ninja_nivel_12"},
		{27, "Teste e Benchmarking", "internal/teste_benchmarking"},
	}

	for _, c := range chapters {
		chapter.Register(&chapter.Chapter{
			Number:  c.number,
			Title:   c.title,
			RootDir: c.rootDir,
		})
	}
}
