package funcoes

import "github.com/fabianoflorentino/aprendago/pkg/format"

func MenuFuncoes([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--sintaxe", ExecFunc: func() { Sintaxe() }},
		{Options: "--desenrolando-enumerando-uma-slice", ExecFunc: func() { DesenrolandoEnumerandoUmaSlice() }},
		{Options: "--defer", ExecFunc: func() { Defer() }},
		{Options: "--metodos", ExecFunc: func() { Metodos() }},
		{Options: "--interfaces-e-polimorfismo", ExecFunc: func() { InterfacesEPolimorfismo() }},
		{Options: "--funcoes-anonimas", ExecFunc: func() { FuncoesAnonimas() }},
		{Options: "--func-como-expressa", ExecFunc: func() { FuncComoExpressa() }},
		{Options: "--retornando-uma-funcao", ExecFunc: func() { RetornandoUmaFuncao() }},
		{Options: "--callback", ExecFunc: func() { Callback() }},
		{Options: "--resolucao-desafio-callback", ExecFunc: func() { ResolucaoDesafioCallback() }},
	}
}
