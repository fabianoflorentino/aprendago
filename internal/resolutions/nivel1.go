package resolutions

import (
	"fmt"
)

func (r Resolution) Nivel1Exercicio1() {
	x := 42
	y := "James Bond"
	z := true

	resolucao := fmt.Sprintf("%v %v %v", x, y, z)

	fmt.Println(resolucao)
}

func (r Resolution) Nivel1Exercicio2() {
	var x int
	var y string
	var z bool

	resolucao := fmt.Sprintf("%v %v %v", x, y, z)

	fmt.Println(resolucao)
}

func (r Resolution) Nivel1Exercicio3() {
	x := 42
	y := "James Bond"
	z := true

	resolucao := fmt.Sprintf("%v %v %v", x, y, z)

	fmt.Println(resolucao)
}

func (r Resolution) Nivel1Exercicio4() {
	type ninja int

	x := ninja(42)

	resolucao := fmt.Sprintf("%v", x)
	fmt.Println(resolucao)

	fmt.Printf("\n%T\n", x)
}

func (r Resolution) Nivel1Exercicio5() {
	type ninja int

	var x ninja
	var y int

	x = 42
	y = int(x)

	resolucao := fmt.Sprintf("%T\n%T", x, y)

	fmt.Println(resolucao)
}

// func (r Resolution) Nivel1Exercicio6() {

// }
