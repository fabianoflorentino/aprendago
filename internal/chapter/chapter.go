// Package chapter provides a collection of functions that represent different
// chapters or sections of a Go programming course. Each function in the package
// corresponds to a specific topic or set of exercises, organized in a sequence
// to facilitate learning. The New function returns a slice of these functions,
// allowing them to be executed in order.
package chapter

import (
	"github.com/fabianoflorentino/aprendago/internal/agrupamento_de_dados"
	"github.com/fabianoflorentino/aprendago/internal/aplicacoes"
	"github.com/fabianoflorentino/aprendago/internal/canais"
	"github.com/fabianoflorentino/aprendago/internal/concorrencia"
	"github.com/fabianoflorentino/aprendago/internal/documentacao"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_1"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_10"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_11"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_12"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_2"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_3"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_4"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_5"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_6"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_7"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_8"
	"github.com/fabianoflorentino/aprendago/internal/exercicios_ninja_nivel_9"
	"github.com/fabianoflorentino/aprendago/internal/fluxo_de_controle"
	"github.com/fabianoflorentino/aprendago/internal/funcoes"
	"github.com/fabianoflorentino/aprendago/internal/fundamentos_da_programacao"
	"github.com/fabianoflorentino/aprendago/internal/ponteiros"
	"github.com/fabianoflorentino/aprendago/internal/seu_ambiente_de_desenvolvimento"
	"github.com/fabianoflorentino/aprendago/internal/structs"
	"github.com/fabianoflorentino/aprendago/internal/teste_benchmarking"
	"github.com/fabianoflorentino/aprendago/internal/tratamento_de_erro"
	"github.com/fabianoflorentino/aprendago/internal/variaveis_valores_tipos"
	"github.com/fabianoflorentino/aprendago/internal/visao_geral_do_curso"
)

// New returns a slice of functions that represent different chapters or topics
// in a Go programming course. Each function in the slice corresponds to a
// specific chapter or exercise level, providing an organized way to access
// various course materials and exercises.
func New() []func() {
	return []func(){
		visao_geral_do_curso.VisaoGeralDoCurso,
		variaveis_valores_tipos.VariaveisValoresETipos,
		exercicios_ninja_nivel_1.ExerciciosNinjaNivel1,
		fundamentos_da_programacao.FundamentosDaProgramacao,
		exercicios_ninja_nivel_2.ExerciciosNinjaNivel2,
		fluxo_de_controle.FluxoDeControle,
		exercicios_ninja_nivel_3.ExerciciosNinjaNivel3,
		agrupamento_de_dados.AgrupamentoDeDados,
		exercicios_ninja_nivel_4.ExerciciosNinjaNivel4,
		structs.TopicStructs,
		exercicios_ninja_nivel_5.ExerciciosNinjaNivel5,
		funcoes.Funcoes,
		exercicios_ninja_nivel_6.ExerciciosNinjaNivel6,
		ponteiros.Ponteiros,
		exercicios_ninja_nivel_7.ExerciciosNinjaNivel7,
		aplicacoes.Aplicacoes,
		exercicios_ninja_nivel_8.ExerciciosNinjaNivel8,
		concorrencia.Concorrencia,
		seu_ambiente_de_desenvolvimento.SeuAmbienteDeDesenvolvimento,
		exercicios_ninja_nivel_9.ExerciciosNinjaNivel9,
		canais.Canais,
		exercicios_ninja_nivel_10.ExerciciosNinjaNivel10,
		tratamento_de_erro.TratamentoDeErro,
		exercicios_ninja_nivel_11.ExerciciosNinjaNivel11,
		documentacao.Documentacao,
		exercicios_ninja_nivel_12.ExerciciosNinjaNivel12,
		teste_benchmarking.TesteEBenchmarking,
	}
}
