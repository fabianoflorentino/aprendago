package exercicios_ninja_nivel_1

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/pkg/format"
)

func NaPraticaExercicio1() {
	topic := format.OutlineContent{
		Title: "Na prática - Exercício #1",
		Content: `
	- Esses são seus primeiros exercícios, e seus primeiros passos.
	- Completando os exercícios dessa seção, você será um ninja nível 1.
	- É o seu primeiro passo pra se tornar um developer ninja.
	- Esses exercícios servem pra reforçar seu aprendizado. Só se aprende a programar programando. Ninguem aprende a andar de bicicleta assistindo vídeos de pessoas andando de bicicleta. É necessário botar a mão na massa.
	- Eu vou começar explicando qual é o exercício. Aí vou pedir pra você dar pausa. Esse é o momento de por os miolos pra trabalhar, encontrar sua solução, tec-tec-tec, e rodar pra ver se funciona.
	Depois é só dar play novamente, ver a minha abordagem para a mesma questão, e comparar nossas soluções.
	- Vamos lá:

	- Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
		1. 42
		2. "James Bond"
		3. true
	- Agora demonstre os valores nestas variáveis, com:
		1. Uma única declaração print
		2. Múltiplas declarações print
	- Solução: https://play.golang.org/p/yYXnWXIQNa ou use no menu do programa (--na-pratica-exercicio-1 --nivel-1 --resolucao)
		`,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio1() {
	x := 42
	y := "James Bond"
	z := true

	resolucao := fmt.Sprintf("%v %v %v", x, y, z)

	format.FormatResolucaoExercicios(resolucao)
}
