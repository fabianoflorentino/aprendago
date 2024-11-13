package exercicios_ninja_nivel_6

import "fmt"

func ResolucaoNaPraticaExercicio1() {
	numero_int := retornaInt()
	numero, texto := retornaIntString()

	fmt.Printf("Inteiro: %v\nInteiro&Texto: %v e %v", numero_int, numero, texto)
}

func ResolucaoNaPraticaExercicio2() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := somaVariadico(slice1...)
	fmt.Printf("Soma do variádico: %v", result)

	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	result2 := somaSlice(slice2)
	fmt.Printf("\nSoma do slice: %v", result2)
}

func ResolucaoNaPraticaExercicio3() {
	defer fmt.Println("Execução do defer ocorreu ao final do contexto.")

	fmt.Println("Execução do defer ocorre ao final do contexto ao qual ela pertence.")
}

func ResolucaoNaPraticaExercicio4() {

	p := pessoa{
		nome:      "Fabiano",
		sobrenome: "Florentino",
		idade:     39,
	}

	p.mostrarPessoa()
}

func ResolucaoNaPraticaExercicio5() {

	q := quadrado{
		lado: 10,
	}

	c := circulo{
		raio: 5,
	}

	fmt.Printf("\nÁrea do quadrado: %f", info(q))
	fmt.Printf("\nÁrea do círculo: %f", info(c))
}

// ResolucaoNaPraticaExercicio1
func retornaInt() int {
	return 42
}

// ResolucaoNaPraticaExercicio1
func retornaIntString() (int, string) {
	return 42, "Olá, mundo!"
}

// ResolucaoNaPraticaExercicio2
func somaVariadico(x ...int) int {
	soma := 0

	for _, inteiro := range x {
		soma += inteiro
	}
	return soma
}

// ResolucaoNaPraticaExercicio2
func somaSlice(slice []int) int {
	soma := 0

	for _, inteiro := range slice {
		soma += inteiro
	}
	return soma
}

// ResolucaoNaPraticaExercicio4
type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

// ResolucaoNaPraticaExercicio4
func (p pessoa) mostrarPessoa() {
	fmt.Printf("Nome: %s %s\nIdade: %d\n", p.nome, p.sobrenome, p.idade)
}

// ResolucaoNaPraticaExercicio5
type figura interface {
	area() float64
}

// ResolucaoNaPraticaExercicio5
func (q quadrado) area() float64 {
	return q.lado * q.lado
}

// ResolucaoNaPraticaExercicio5
func (c circulo) area() float64 {
	return 2 * 3.14 * c.raio
}

// ResolucaoNaPraticaExercicio5
type quadrado struct {
	lado float64
}

// ResolucaoNaPraticaExercicio5
type circulo struct {
	raio float64
}

// ResolucaoNaPraticaExercicio5
func info(f figura) float64 {
	return f.area()
}
