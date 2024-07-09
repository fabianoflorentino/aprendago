package fundamentos_da_programacao

import "github.com/fabianoflorentino/aprendago/pkg/format"

func MenuFundamentosDaProgramcao([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--tipo-booleano", ExecFunc: func() { TipoBooleano() }},
		{Options: "--como-os-computadores-funcionam", ExecFunc: func() { ComoOsComputadoresFuncionam() }},
		{Options: "--tipos-numericos", ExecFunc: func() { TiposNumericos() }},
		{Options: "--overflow", ExecFunc: func() { Overflow() }},
		{Options: "--tipo-string", ExecFunc: func() { TipoString() }},
		{Options: "--sistemas-numericos", ExecFunc: func() { SistemasNumericos() }},
		{Options: "--constantes", ExecFunc: func() { Constantes() }},
		{Options: "--iota", ExecFunc: func() { Iota() }},
		{Options: "--deslocamento-de-bits", ExecFunc: func() { DeslocamentoDeBits() }},
	}
}
