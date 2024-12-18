package exercicios_ninja_nivel_11

import (
	"encoding/json"
	"fmt"
)

// person represents an individual with a first name, last name, and a list of sayings.
type person struct {
	First   string
	Last    string
	Sayings []string
}

func ResolucaoNaPraticaExercicio1() {
	// - Utilizando este código: https://play.golang.org/p/3W69TH4nON
	// - ...remova o underscore e verifique e lide com o erro de maneira apropriada.
	// - Solução: https://github.com/ellenkorbes/aprendago/blob/master/c%C3%B3digo/24_exercicios-ninja-11/01/main.go
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
}

func ResolucaoNaPraticaExercicio2() {}

func ResolucaoNaPraticaExercicio3() {}

func ResolucaoNaPraticaExercicio4() {}

func ResolucaoNaPraticaExercicio5() {}
