package format

import (
	"errors"
	"fmt"
	"strings"
)

// CreateQuestionnaire is an interface for creating a questionnaire with questions and answer options.
type CreateQuestionnaire interface {

	// Create generates a questionnaire from a slice of Questionnaire.
	// It returns a slice of Questionnaire.
	Create(Questionnaire []Questionnaire) []Questionnaire
}

// CollectAnswers is an interface for collecting answers from a questionnaire.
type CollectAnswers interface {

	// Collect gathers answers from a questionnaire and returns a slice of Answer and an error.
	// It receives a slice of Questionnaire and returns a slice of Answer and an error.
	Collect(Questionnaire []Questionnaire) ([]Answer, error)
}

// ValidateAnswer is an interface for validating answers from a questionnaire.
type ValidateAnswer interface {

	// Validate checks the answers from a questionnaire and displays the test result.
	// It receives a slice of Answer and a slice of Answer.
	Validate(answer []Answer, key []Answer)
}

// Questionnaire represents a question in a test with its answer options.
type Questionnaire struct {
	Number   string // The number of the question
	Question string // The text of the question
	Options  string // The available answer options
}

// Answer represents the answer to a question, storing the user's response and its status.
type Answer struct {
	Number string // The number of the question
	Answer string // The answer given by the user
	Status string // The status of the answer (e.g., correct, incorrect)
}

// QuestionnaireService is a type that implements the CreateQuestionnaire interface through the Create method.
type QuestionnaireService struct{}

// Create generates a questionnaire from a slice of Questionnaire.
// It receives a slice of Questionnaire and returns a slice of Questionnaire.
func (qs QuestionnaireService) Create(questionnaire []Questionnaire) []Questionnaire {
	var test []Questionnaire

	// Iterate over the slice of Questionnaire and add the questions and answer options to the test slice.
	for _, quest := range questionnaire {
		test = append(test, Questionnaire{Number: quest.Number, Question: quest.Question, Options: quest.Options})
	}

	return test
}

// AnswerService is a type that implements the CollectAnswers interface through the Collect method.
type AnswerService struct{}

// Collect gathers answers from a questionnaire and returns a slice of Answer and an error.
// It receives a slice of Questionnaire and returns a slice of Answer and an error.
func (as AnswerService) Collect(questionnaire []Questionnaire) ([]Answer, error) {
	var answerList []Answer
	var answer string

	// Create a questionnaire with questions and answer options.
	create_questionnaire := new(QuestionnaireService).Create(questionnaire)

	// Iterate over the slice of Questionnaire and collect the user's answers for each question.
	for _, q := range create_questionnaire {
		fmt.Print(q.Number, q.Question, q.Options)
		fmt.Scan(&answer)

		// Check if the answer is valid.
		if !validateOption(answer, q.Options) {
			return nil, errors.New("invalid option")
		}

		// Add the answer to the slice of Answer.
		answerList = append(answerList, Answer{q.Number, answer, ""})
	}

	return answerList, nil
}

// Validate checks the answers from a questionnaire and displays the test result.
// It receives a slice of Answer and a slice of Answer.
func (as AnswerService) Validate(answer []Answer, key []Answer) {
	var validAnswer []Answer
	var correctAnswerCount int

	// Iterate over the user's answers and check if the answer is correct or incorrect.
	// Add the answer to the slice of ValidAnswer.
	for i, ans := range answer {
		if ans.Answer == key[i].Answer {
			validAnswer = append(validAnswer, Answer{Number: ans.Number, Answer: ans.Answer, Status: "correct"})
			correctAnswerCount++
		} else {
			validAnswer = append(validAnswer, Answer{Number: ans.Number, Answer: ans.Answer, Status: "incorrect"})
		}
	}

	fmt.Printf("\nTest result:\n\n")
	fmt.Printf("Total correct answers %d out of %d\n", correctAnswerCount, len(key))

	// Calculate the user's test score percentage.
	testResult := float64(correctAnswerCount) / float64(len(key)) * 100
	fmt.Printf("Score percentage %.2f%%\n\n", testResult)

	// Display the user's answers and the status of each answer.
	for _, ans := range validAnswer {

		// Remove the period from the question number.
		ansNumber := strings.Trim(ans.Number, ".")

		fmt.Printf("Question %s: Answer: %s - %s\n", ansNumber, ans.Answer, ans.Status)
	}
}

// validateOption checks if the user's answer is valid for a question.
// It receives the user's answer and the valid options as strings.
// It returns true if the answer is among the valid options, otherwise it returns false.
func validateOption(answer, options string) bool {
	// Iterate over the answer options of the question and check if the user's answer is valid.
	for _, opt := range strings.Fields(options) {
		opt = strings.Trim(opt, "[]")
		opt = strings.Split(opt, " ")[0] // Extract the option letter
		if answer == opt {
			return true
		}
	}
	return false
}
