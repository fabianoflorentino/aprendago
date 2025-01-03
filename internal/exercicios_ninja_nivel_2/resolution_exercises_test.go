// Package exercicios_ninja_nivel_2 contains exercises for practicing Go programming at a beginner level.
// It includes various functions and tests to help users understand and apply basic concepts of the Go language.
package exercicios_ninja_nivel_2

import (
	"fmt"
	"testing"
)

// TestResolucaoNaPraticaExercicio1 tests the formatting of a number in decimal, binary, and hexadecimal formats.
// It verifies that the formatted string matches the expected output.
func TestResolucaoNaPraticaExercicio1(t *testing.T) {
	numero := 1000019

	resolucao := fmt.Sprintf("Decimal: %d \nBinário: %b \nHexadecimal: %#x \n", numero, numero, numero)

	expect := "Decimal: 1000019 \nBinário: 11110100001001010011 \nHexadecimal: 0xf4253 \n"
	if resolucao != expect {
		t.Errorf("\nExpected: \n%v\nGot:\n%v", expect, resolucao)
	}
}

// TestResolucaoNaPraticaExercicio2 tests the comparison operations between two integer values.
// It verifies that the formatted string matches the expected output.
func TestResolucaoNaPraticaExercicio2(t *testing.T) {
	a := 10
	b := 20

	resolucao := fmt.Sprintf("a: %v \nb: %v \n\na == b: %v \na != b: %v \na <= b: %v \na < b: %v \na >= b: %v \na > b: %v \n", a, b, a == b, a != b, a <= b, a < b, a >= b, a > b)

	expect := "a: 10 \nb: 20 \n\na == b: false \na != b: true \na <= b: true \na < b: true \na >= b: false \na > b: false \n"
	if resolucao != expect {
		t.Errorf("\nExpected: \n%v\nGot:\n%v", expect, resolucao)
	}
}

// TestResolucaoNaPraticaExercicio3 tests the use of untyped and typed constants in Go.
// It verifies that the formatted string matches the expected output.
func TestResolucaoNaPraticaExercicio3(t *testing.T) {
	const (
		untypedConst     = 10
		typedConst   int = 20
	)

	resolucao := fmt.Sprintf("untypedConst: %v \ntypedConst: %v", untypedConst, typedConst)

	expect := "untypedConst: 10 \ntypedConst: 20"
	if resolucao != expect {
		t.Errorf("\nExpected: \n%v\nGot:\n%v", expect, resolucao)
	}
}

// TestResolucaoNaPraticaExercicio4 tests the formatting of an integer value in decimal, binary, and hexadecimal formats.
// It verifies the correctness of the formatted strings for the initial value and its left-shifted result.
func TestResolucaoNaPraticaExercicio4(t *testing.T) {
	v1 := 42

	resolucao1 := fmt.Sprintf("Decimal: %d \nBinário: %b \nHexadecimal: %#x \n", v1, v1, v1)

	v2 := v1 << 1

	resolucao2 := fmt.Sprintf("Decimal: %d \nBinário: %b \nHexadecimal: %#x \n", v2, v2, v2)

	expect1 := "Decimal: 42 \nBinário: 101010 \nHexadecimal: 0x2a \n"
	expect2 := "Decimal: 84 \nBinário: 1010100 \nHexadecimal: 0x54 \n"

	if resolucao1 != expect1 {
		t.Errorf("\nExpected: \n%v\nGot:\n%v", expect1, resolucao1)
	}

	if resolucao2 != expect2 {
		t.Errorf("\nExpected: \n%v\nGot:\n%v", expect2, resolucao2)
	}
}

// TestResolucaoNaPraticaExercicio6 tests the generation of consecutive years
// starting from 2024 using the iota enumerator. It verifies that the generated
// string matches the expected format and values for the years 2024, 2025, 2026,
// and 2027.
func TestResolucaoNaPraticaExercicio6(t *testing.T) {
	const (
		currentYear = 2024 + iota
		nextYear
		nextNextYear
		nextNextNextYear
	)

	resolucao := fmt.Sprintf("currentYear: %v \nnextYear: %v \nnextNextYear: %v \nnextNextNextYear: %v", currentYear, nextYear, nextNextYear, nextNextNextYear)
	expect := "currentYear: 2024 \nnextYear: 2025 \nnextNextYear: 2026 \nnextNextNextYear: 2027"

	if resolucao != expect {
		t.Errorf("\nExpected: \n%v\nGot:\n%v", expect, resolucao)
	}
}
