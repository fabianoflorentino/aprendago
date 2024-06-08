package outline

import (
	"fmt"

	format "github.com/fabianoflorentino/aprendago/outline/format"
)

func NaPraticaExercicio10() {
	topic := format.OutlineContent{
		Title: "Na prática: Exercício #10",
		Content: `
- Anote (à mão) o resultado das expressões:
  - fmt.Println(true && true)
  - fmt.Println(true && false)
  - fmt.Println(true || true)
  - fmt.Println(true || false)
  - fmt.Println(!true)
- Ninja nível 3! Parabéns!
    `,
	}

	format.FormatOutlineTopic(topic)
}

func ResolucaoNaPraticaExercicio10() {

	fmt.Println("Anotado!")
	anotacao := `
true && true
true && false
true || true
true || false
!true
  `

	fmt.Printf("%v", anotacao)
}
