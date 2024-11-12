package exercicios_ninja_nivel_5

import "fmt"

func ResolucaoNaPraticaExercicio1() {
	type pessoa struct {
		nome               string
		sobrenome          string
		sabores_de_sorvete []string
	}

	p1 := pessoa{
		nome:               "Fulano",
		sobrenome:          "de Tal",
		sabores_de_sorvete: []string{"Chocolate", "Morango", "Baunilha"},
	}

	p2 := pessoa{
		nome:               "Ciclano",
		sobrenome:          "da Silva",
		sabores_de_sorvete: []string{"Pistache", "Creme", "Coco"},
	}

	pessoas := []pessoa{p1, p2}

	for _, p := range pessoas {
		fmt.Println(p.nome, p.sobrenome)
		for _, s := range p.sabores_de_sorvete {
			fmt.Printf("\t%v\n", s)
		}
	}
}

func ResolucaoNaPraticaExercicio2() {
	type pessoa struct {
		nome               string
		sobrenome          string
		sabores_de_sorvete []string
	}

	p1 := pessoa{
		nome:               "Fulano",
		sobrenome:          "de Tal",
		sabores_de_sorvete: []string{"Chocolate", "Morango", "Baunilha"},
	}

	p2 := pessoa{
		nome:               "Ciclano",
		sobrenome:          "da Silva",
		sabores_de_sorvete: []string{"Pistache", "Creme", "Coco"},
	}

	pessoas := map[string]pessoa{
		p1.sobrenome: p1,
		p2.sobrenome: p2,
	}

	for k, p := range pessoas {
		println(k)
		for _, s := range p.sabores_de_sorvete {
			println("\t", s)
		}
	}
}

func ResolucaoNaPraticaExercicio3() {
	type veiculo struct {
		portas int
		cor    string
	}

	type caminhonete struct {
		veiculo
		tracaoNasQuatro bool
	}

	type sedan struct {
		veiculo
		modeloLuxo bool
	}

	cnt := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor:    "Preto",
		},

		tracaoNasQuatro: true,
	}

	sdn := sedan{
		veiculo: veiculo{
			portas: 2,
			cor:    "Branco",
		},

		modeloLuxo: true,
	}

	fmt.Printf("Caminhonete: %+v \n", cnt)
	fmt.Printf("Sedan: %+v \n", sdn)

	fmt.Println("Portas da caminhonete:", cnt.portas)
	fmt.Println("Cor do sedan:", sdn.cor)
}

func ResolucaoNaPraticaExercicio4() {
	pessoa := struct {
		Nome             string
		Idade            int
		Amigos           map[string]int
		ComidasFavoritas []string
	}{
		Nome:  "Fabiano",
		Idade: 39,
		Amigos: map[string]int{
			"Ale":   46,
			"Lucas": 38,
		},
		ComidasFavoritas: []string{
			"Pizza",
			"Lasanha",
			"Hamburguer",
		},
	}

	// mais para frente
	// jsonData, _ := json.Marshal(pessoa)
	// fmt.Println(string(jsonData))

	fmt.Println(pessoa)
}
