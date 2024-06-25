package exercicios_ninja_nivel_5

import "github.com/fabianoflorentino/aprendago/outline/format"

func MenuExerciciosNinjaNivel5([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--na-pratica-exercicio-1 --nivel-5", ExecFunc: func() { NaPraticaExercicio1() }},
		{Options: "--na-pratica-exercicio-1 --nivel-5 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--na-pratica-exercicio-2 --nivel-5", ExecFunc: func() { NaPraticaExercicio2() }},
		{Options: "--na-pratica-exercicio-2 --nivel-5 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--na-pratica-exercicio-3 --nivel-5", ExecFunc: func() { NaPraticaExercicio3() }},
		{Options: "--na-pratica-exercicio-3 --nivel-5 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--na-pratica-exercicio-4 --nivel-5", ExecFunc: func() { NaPraticaExercicio4() }},
		{Options: "--na-pratica-exercicio-4 --nivel-5 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
	}
}
