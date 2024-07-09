package fluxo_de_controle

import "github.com/fabianoflorentino/aprendago/pkg/format"

func MenuFluxoDeControle([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--entendendo-fluxo-de-controle", ExecFunc: func() { EntendendoFluxoDeControle() }},
		{Options: "--loops-inicializacao-condicao-pos", ExecFunc: func() { LoopsInicializacaoCondicaoPos() }},
		{Options: "--loops-nested-loop", ExecFunc: func() { LoopsNestedLoop() }},
		{Options: "--loops-a-declaracao-for", ExecFunc: func() { LoopsADeclaracaoFor() }},
		{Options: "--loops-break-e-continue", ExecFunc: func() { LoopsBreakEContinue() }},
		{Options: "--loops-utilizando-ascii", ExecFunc: func() { LoopsUtilizandoAscii() }},
		{Options: "--loops-utilizando-ascii --resolucao", ExecFunc: func() { ResolucaoLoopsUtilizandoAscii() }},
		{Options: "--condicionais-a-declaracao-if", ExecFunc: func() { CondicionaisADeclaracaoIf() }},
		{Options: "--condicionais-if-else-if-else", ExecFunc: func() { CondicionaisIfElseIfElse() }},
		{Options: "--condicionais-a-declaracao-switch", ExecFunc: func() { CondicionaisADeclaracaoSwitch() }},
		{Options: "--condicionais-a-declaracao-switch-pt2", ExecFunc: func() { CondicionaisADeclaracaoSwitchPt2() }},
		{Options: "--operadores-logicos-condicionais", ExecFunc: func() { OperadoresLogicosCondicionais() }},
	}
}
