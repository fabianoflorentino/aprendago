package variaveis_valores_tipos

import "github.com/fabianoflorentino/aprendago/pkg/format"

func MenuVariaveisValoresTipos([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--go-playground", ExecFunc: func() { GoPlayground() }},
		{Options: "--hello-world", ExecFunc: func() { HelloWorld() }},
		{Options: "--operador-curto-de-declaracao", ExecFunc: func() { OperadorCurtoDeDeclaracao() }},
		{Options: "--a-palavra-reservada-var", ExecFunc: func() { ApalavraReservadaVar() }},
		{Options: "--explorando-tipos", ExecFunc: func() { ExplorandoTipos() }},
		{Options: "--valor-zero", ExecFunc: func() { ValorZero() }},
		{Options: "--o-pacote-fmt", ExecFunc: func() { OpacoteFmt() }},
		{Options: "--criando-seu-proprio-tipo", ExecFunc: func() { CriandoSeuProprioTipo() }},
		{Options: "--conversao-nao-coercao", ExecFunc: func() { ConversaoNaoCoercao() }},
	}
}
