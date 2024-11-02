package visao_geral_do_curso

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func VisaoGeralDoCurso() {
	fmt.Printf("01 - Visão Geral do Curso\n\n")

	BemVindo()
	PorQueGo()
	Sucesso()
	Recursos()
	ComoEsseCursoFunciona()
}

func BemVindo() {
	format.BuildSection(rootDir, "Bem-vindo!")
}

func PorQueGo() {
	format.BuildSection(rootDir, "Por que Go?")
}

func Sucesso() {
	format.BuildSection(rootDir, "Sucesso")
}

func Recursos() {
	format.BuildSection(rootDir, "Recursos")
}

func ComoEsseCursoFunciona() {
	format.BuildSection(rootDir, "Como esse curso funciona")
}
