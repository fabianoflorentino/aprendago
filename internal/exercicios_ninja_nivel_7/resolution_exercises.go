package exercicios_ninja_nivel_7

import "fmt"

func ResolucaoNaPraticaExercicio1() {
	valor := "valor"
	println(&valor)
}

func ResolucaoNaPraticaExercicio2() {
	p := pessoa{
		nome:      "Fulano",
		sobrenome: "de Tal",
		idade:     25,
	}

	fmt.Printf("Original -> Nome: %s %s, Idade: %d\n", p.nome, p.sobrenome, p.idade)

	mudeMe(&p)

	fmt.Printf("Alterado -> Nome: %s %s, Idade: %d\n", p.nome, p.sobrenome, p.idade)
}

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func mudeMe(p *pessoa) {
	p.nome = "João"
	p.sobrenome = "Silva"
	p.idade = 30
}
