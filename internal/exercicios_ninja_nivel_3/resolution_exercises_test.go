package exercicios_ninja_nivel_3

import (
	"fmt"
	"testing"
)

func TestResolucaoNaPraticaExercicio1(t *testing.T) {
	var result string

	for idx := 0; idx < 5; idx++ {
		result += fmt.Sprintf("%v ", idx+1)
	}

	expect := "1 2 3 4 5 "

	if result != expect {
		t.Errorf("\ngot %v\nwant %v", result, expect)
	}
}
