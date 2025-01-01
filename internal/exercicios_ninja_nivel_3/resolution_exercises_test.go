// Package exercicios_ninja_nivel_3 contains solutions and tests for exercises
// designed to practice and improve Go programming skills at an intermediate level.
// These exercises cover various topics such as loops, conditionals, and string manipulation.
package exercicios_ninja_nivel_3

import (
	"fmt"
	"strings"
	"testing"
	"time"
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

// TestResolucaoExercicioNaPratica3 is a test function that calculates and prints the years from a given birth year (1985) to the current year.
// It also calculates the age based on the current year and compares the result with an expected string.
// If the calculated age string does not match the expected string, it prints the actual and expected values.
func TestResolucaoExercicioNaPratica3(t *testing.T) {
	var yearsCount string
	var birthYear int = 1985

	currentYear := time.Now().Year()

	for year := birthYear; year <= currentYear; year++ {
		if year == currentYear {
			yearsCount += fmt.Sprintf("%d\n", year)
		} else {
			yearsCount += fmt.Sprintf("%d, ", year)
		}
	}

	age := fmt.Sprintf("%v\nSua idade: %v", yearsCount, currentYear-birthYear)

	expect := `
1985, 1986, 1987, 1988, 1989, 1990, 1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998, 1999, 2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010, 2011, 2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020, 2021, 2022, 2023, 2024

Sua idade: 39
`

	if strings.TrimSpace(age) != strings.TrimSpace(expect) {
		t.Errorf("got\n%v\nwant\n%v", age, expect)
	}
}

// TestResolucaoExercicioNaPratica4 tests the functionality of generating a string
// that lists all years from a given birth year (1985) up to the current year,
// followed by the calculated age based on the current year. It compares the
// generated string with an expected string to ensure correctness.
func TestResolucaoExercicioNaPratica4(t *testing.T) {
	var yearsCount string
	var birthYear int = 1985

	currentYear := time.Now().Year()

	for {
		if birthYear >= currentYear {
			yearsCount += fmt.Sprintf("%d", birthYear)
			break
		} else {
			yearsCount += fmt.Sprintf("%d, ", birthYear)
			birthYear++
		}
	}

	age := fmt.Sprintf("%v\nSua idade: %v", yearsCount, time.Now().Year()-1985)

	expect := `
1985, 1986, 1987, 1988, 1989, 1990, 1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998, 1999, 2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010, 2011, 2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020, 2021, 2022, 2023, 2024
Sua idade: 39
`

	if strings.TrimSpace(age) != strings.TrimSpace(expect) {
		t.Errorf("\ngot: %v\nwant: %v", age, expect)
	}
}

// TestResolucaoNaPraticaExercicio5 tests the output of a loop that generates a formatted string
// of numbers from 10 to 100 and their remainders when divided by 4. The output is formatted
// such that every 10 numbers are separated by a newline. The test compares the generated
// string with the expected string to ensure correctness.
func TestResolucaoNaPraticaExercicio5(t *testing.T) {
	var result string
	result += "\n"

	for number := 10; number <= 100; number++ {
		if number%10 == 0 && number != 10 {
			result = strings.TrimSuffix(result, ", ") + ",\n"
		}

		result += fmt.Sprintf("%d: %d, ", number, number%4)
	}

	result = strings.TrimSuffix(result, ", ") + "\n"

	expect := `
10: 2, 11: 3, 12: 0, 13: 1, 14: 2, 15: 3, 16: 0, 17: 1, 18: 2, 19: 3,
20: 0, 21: 1, 22: 2, 23: 3, 24: 0, 25: 1, 26: 2, 27: 3, 28: 0, 29: 1,
30: 2, 31: 3, 32: 0, 33: 1, 34: 2, 35: 3, 36: 0, 37: 1, 38: 2, 39: 3,
40: 0, 41: 1, 42: 2, 43: 3, 44: 0, 45: 1, 46: 2, 47: 3, 48: 0, 49: 1,
50: 2, 51: 3, 52: 0, 53: 1, 54: 2, 55: 3, 56: 0, 57: 1, 58: 2, 59: 3,
60: 0, 61: 1, 62: 2, 63: 3, 64: 0, 65: 1, 66: 2, 67: 3, 68: 0, 69: 1,
70: 2, 71: 3, 72: 0, 73: 1, 74: 2, 75: 3, 76: 0, 77: 1, 78: 2, 79: 3,
80: 0, 81: 1, 82: 2, 83: 3, 84: 0, 85: 1, 86: 2, 87: 3, 88: 0, 89: 1,
90: 2, 91: 3, 92: 0, 93: 1, 94: 2, 95: 3, 96: 0, 97: 1, 98: 2, 99: 3,
100: 0
`

	normalize := func(s string) string {
		return strings.ReplaceAll(strings.TrimSpace(s), "\r\n", "\n")
	}

	if normalize(result) != normalize(expect) {
		t.Errorf("\ngot:\n%v\nwant:\n%v", result, expect)
	}
}
