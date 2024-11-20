# [Questionnaire](../pkg/format/questionnaire.go)

This Go package creates, collects, and validates a questionnaire. It defined some interfaces and structures to manage the questionnaire and answers.

## Interfaces

### [1. CreateQuestionarie](../pkg/format/questionnaire.go#L10)

- Is defined as the method `Create` that will create a questionnaire from a slice of the questionnaire.
- Returns a Slice of the Questionnaire.

### [2. CollectAnswers](../pkg/format/questionnaire.go#L18)

- Is defined as the method `Collect` that will collect the answers from a questionnaire.
- Receive a slice of a questionnaire and return a slice of Answers and an error.

### [3. ValidateAnswers](../pkg/format/questionnaire.go#L26)

- Is defined as the method `Validate` that will validate the answers from a questionnaire.
- Receive two slices of Answers (user answers and correct answers).

## Structures

### [1. Questionnaire](../pkg/format/questionnaire.go#L34)

- Represents a question in a test with options for answers
- Fields:
  - `Number`: The number of the question.
  - `Question`: The question.
  - `Options`: The options of the question.

### [2. Answer](../pkg/format/questionnaire.go#L41)

- Represents an answer to a question, storing the answer of the user and status.
- Fields:
  - `Number`: The number of the question.
  - `Answer`: The answer of the user.
  - `Status`: The status of the answer (correct or incorrect).

## [3. QuestionnaireService](../pkg/format/questionnaire.go#L48)

- Implements the interface `CreateQuestionnaire` through the method `Create`.
- Method `Create`:
  - Receives a slice of questions and returns a slice of the questionnaire.
  - Iterated over the slice of the questionnaire and added a question and options to the slice test.

### [4. AnswerService](../pkg/format/questionnaire.go#L64)

- Implements the interface `CollectAnswers` and `ValidateAnswers` through the methods `Collect` and `Validate`.
- [Method `Collect`](../pkg/format/questionnaire.go#L68):
  - Receives a slice of the questionnaire and returns a slice of answers and an error.
  - Create a questionnaire using `QuestionnaireService`.
  - Iterated the questionnaire and collected the answers from the user.
  - Validate the answers using the function `validateOption`.
- [Method `Validate`](../pkg/format/questionnaire.go#L94):
  - Receives two slices of answers (user answers and correct answers).
  - Iterated the answers and validated the answers checking if the answer was correct or incorrect.
  - Calculate the score of the user and return the score.

## Auxiliary Functions

### [1. validateOption](../pkg/format/questionnaire.go#L129)

- Checking if the user's answer is valid to answer the question.
- Receives an answer from the user and validates if the answer is a string
- Returns `true` if the answer is between the valid options otherwise returns `false`.

## How it works

### 1. Creates Questionnaire

- The questionnaire is created when using the method `Create` from `QuestionnaireService`

### 2. Collect the Answers

- The answers are collected when using the method `Collect` from `AnswerService`
- The answers are validated using the function `validateOption`

### 3. Validate the Answers

- The answers are validated when using the method `Validate` from `AnswerService`
- The results show up including the score count of answers and correct percentage.
