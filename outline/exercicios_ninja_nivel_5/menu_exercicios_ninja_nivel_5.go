package exercicios_ninja_nivel_5

import "github.com/fabianoflorentino/aprendago/outline/format"

func MenuExerciciosNinjaNivel5([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--na-pratica-exercicio-1 --nivel-5", ExecFunc: func() { NaPraticaExercicio1() }},
		{Options: "--na-pratica-exercicio-1 --nivel-5 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
	}
}
