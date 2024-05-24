// https://github.com/vkorbes/aprendago/blob/master/OUTLINE.md#na-pr%C3%A1tica-exerc%C3%ADcio-3
package exercicio

import "fmt"

func Exercicio03() {
	x := 42
	y := "James Bond"
	z := true

	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}
