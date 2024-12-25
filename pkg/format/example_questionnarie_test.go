// Package format provides functionality for creating and managing questionnaires.
// It includes services for generating questionnaires and collecting answers.
package format

import (
	"fmt"
)

// ExampleCreate demonstrates how to create a questionnaire with questions and answer options,
// generate a questionnaire using the QuestionnaireService, and print the created questionnaire.
func ExampleCreate() {
	// Create a questionnaire with questions and answer options
	questionnaire := []Questionnaire{
		{Number: "1", Question: "What is the capital of France?", Options: "a) Paris b) London c) Rome"},
		{Number: "2", Question: "What is the capital of Italy?", Options: "a) Paris b) London c) Rome"},
	}

	// Create a new questionnaire service
	qs := QuestionnaireService{}

	// Generate a questionnaire from the slice of Questionnaire
	test := qs.Create(questionnaire)

	// Print the generated questionnaire
	for _, q := range test {
		fmt.Printf("%s. %s\n", q.Number, q.Question)
		fmt.Println(q.Options)
	}

	// Output:
	// 1. What is the capital of France?
	// a) Paris b) London c) Rome
	// 2. What is the capital of Italy?
	// a) Paris b) London c) Rome
}

// ExampleCollect demonstrates how to create a questionnaire, generate it using a
// QuestionnaireService, and collect answers using an AnswerService. It prints the
// collected answers or an error message if the collection fails.
func ExampleCollect() {}

// ExampleValidate demonstrates how to use the AnswerService to validate user answers
// against a key of correct answers. It creates a key with correct answers and a list
// of user answers, then validates the user answers using the AnswerService. The
// results are printed to the console, showing which answers are correct and which
// are incorrect, followed by the expected output for comparison.
func ExampleValidate() {
	// Create a key with correct answers
	key := []Answer{
		{Number: "1", Answer: "a"},
		{Number: "2", Answer: "c"},
	}

	// Create a list of user answers
	answers := []Answer{
		{Number: "1", Answer: "a"},
		{Number: "2", Answer: "b"},
	}

	// Create a new answer service
	as := AnswerService{}

	// Validate the user answers against the key
	as.Validate(answers, key)

	for _, ans := range answers {
		fmt.Printf("%s. %s - %s\n", ans.Number, ans.Answer, ans.Status)
	}

	expect := []string{
		"1. a - correct",
		"2. b - incorrect",
	}

	fmt.Println("Expected output:")
	for _, e := range expect {
		fmt.Println(e)
	}
}
