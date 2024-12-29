// Package exercicios_ninja_nivel_3 contains solutions and tests for exercises
// designed to practice and improve Go programming skills at an intermediate level.
// These exercises cover various topics such as loops, conditionals, and string manipulation.
package exercicios_ninja_nivel_3

import (
	"fmt"
	"testing"
)

// TestResolucaoNaPraticaExercicio1 tests the functionality of a loop that
// concatenates numbers from 1 to 5 into a single string with spaces in between.
// It compares the result with the expected string "1 2 3 4 5 " and reports an error if they do not match.
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

// TestResolucaoNaPraticaExercicio2 tests the output of a loop that generates
// a formatted string of uppercase ASCII letters (A-Z) and their corresponding
// Unicode code points. Each letter is repeated three times with different
// Unicode code points (U+0000, U+0001, U+0002). The test compares the generated
// string with the expected string and reports an error if they do not match.
func TestResolucaoNaPraticaExercicio2(t *testing.T) {
	var result string

	for upperLetter := 65; upperLetter <= 90; upperLetter++ {
		result += fmt.Sprintf("%d: ", upperLetter)

		for threeTimes := 0; threeTimes < 3; threeTimes++ {
			result += fmt.Sprintf(" %#U '%c'", threeTimes, upperLetter)
		}
		result += "\n"
	}

	expect := `65:  U+0000 'A' U+0001 'A' U+0002 'A'
66:  U+0000 'B' U+0001 'B' U+0002 'B'
67:  U+0000 'C' U+0001 'C' U+0002 'C'
68:  U+0000 'D' U+0001 'D' U+0002 'D'
69:  U+0000 'E' U+0001 'E' U+0002 'E'
70:  U+0000 'F' U+0001 'F' U+0002 'F'
71:  U+0000 'G' U+0001 'G' U+0002 'G'
72:  U+0000 'H' U+0001 'H' U+0002 'H'
73:  U+0000 'I' U+0001 'I' U+0002 'I'
74:  U+0000 'J' U+0001 'J' U+0002 'J'
75:  U+0000 'K' U+0001 'K' U+0002 'K'
76:  U+0000 'L' U+0001 'L' U+0002 'L'
77:  U+0000 'M' U+0001 'M' U+0002 'M'
78:  U+0000 'N' U+0001 'N' U+0002 'N'
79:  U+0000 'O' U+0001 'O' U+0002 'O'
80:  U+0000 'P' U+0001 'P' U+0002 'P'
81:  U+0000 'Q' U+0001 'Q' U+0002 'Q'
82:  U+0000 'R' U+0001 'R' U+0002 'R'
83:  U+0000 'S' U+0001 'S' U+0002 'S'
84:  U+0000 'T' U+0001 'T' U+0002 'T'
85:  U+0000 'U' U+0001 'U' U+0002 'U'
86:  U+0000 'V' U+0001 'V' U+0002 'V'
87:  U+0000 'W' U+0001 'W' U+0002 'W'
88:  U+0000 'X' U+0001 'X' U+0002 'X'
89:  U+0000 'Y' U+0001 'Y' U+0002 'Y'
90:  U+0000 'Z' U+0001 'Z' U+0002 'Z'
`

	if result != expect {
		t.Errorf("\ngot %v\nwant %v", result, expect)
	}
}
