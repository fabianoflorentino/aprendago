package format

import (
	"testing"
)

// TestValidateCorrectAnswers tests the Validate function with all correct answers.
func TestValidateCorrectAnswers(t *testing.T) {
	answerService := AnswerService{}
	answers := []Answer{
		{Number: "1", Answer: "A"},
		{Number: "2", Answer: "B"},
		{Number: "3", Answer: "C"},
	}
	key := []Answer{
		{Number: "1", Answer: "A"},
		{Number: "2", Answer: "B"},
		{Number: "3", Answer: "C"},
	}

	answerService.Validate(answers, key)
}

// TestValidateIncorrectAnswers tests the Validate function with all incorrect answers.
func TestValidateIncorrectAnswers(t *testing.T) {
	answerService := AnswerService{}
	answers := []Answer{
		{Number: "1", Answer: "B"},
		{Number: "2", Answer: "C"},
		{Number: "3", Answer: "A"},
	}
	key := []Answer{
		{Number: "1", Answer: "A"},
		{Number: "2", Answer: "B"},
		{Number: "3", Answer: "C"},
	}

	answerService.Validate(answers, key)
}

// TestValidateMixedAnswers tests the Validate function with a mix of correct and incorrect answers.
func TestValidateMixedAnswers(t *testing.T) {
	answerService := AnswerService{}
	answers := []Answer{
		{Number: "1", Answer: "A"},
		{Number: "2", Answer: "C"},
		{Number: "3", Answer: "C"},
	}
	key := []Answer{
		{Number: "1", Answer: "A"},
		{Number: "2", Answer: "B"},
		{Number: "3", Answer: "C"},
	}

	answerService.Validate(answers, key)
}

// TestValidateOptionValid tests the validateOption function with valid options.
func TestValidateOptionValid(t *testing.T) {
	options := "[A] Option1 [B] Option2 [C] Option3"
	if !validateOption("A", options) {
		t.Errorf("Expected true, got false")
	}
	if !validateOption("B", options) {
		t.Errorf("Expected true, got false")
	}
	if !validateOption("C", options) {
		t.Errorf("Expected true, got false")
	}
}

// TestValidateOptionInvalid tests the validateOption function with invalid options.
func TestValidateOptionInvalid(t *testing.T) {
	options := "[A] Option1 [B] Option2 [C] Option3"
	if validateOption("D", options) {
		t.Errorf("Expected false, got true")
	}
	if validateOption("E", options) {
		t.Errorf("Expected false, got true")
	}
	if validateOption("F", options) {
		t.Errorf("Expected false, got true")
	}
}

// TestValidateOptionEmpty tests the validateOption function with empty options.
func TestValidateOptionEmpty(t *testing.T) {
	options := ""
	if validateOption("A", options) {
		t.Errorf("Expected false, got true")
	}
	if validateOption("B", options) {
		t.Errorf("Expected false, got true")
	}
	if validateOption("C", options) {
		t.Errorf("Expected false, got true")
	}
}
