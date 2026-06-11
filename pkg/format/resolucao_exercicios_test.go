package format

import (
	"strings"
	"testing"
)

func TestFormatResolucaoExercicios(t *testing.T) {
	buf, cleanup := captureStdout(t)
	FormatResolucaoExercicios("some resolution text")
	cleanup()

	output := buf.String()
	if !strings.Contains(output, "Resolução:") {
		t.Error("output missing 'Resolução:' header")
	}
	if !strings.Contains(output, "some resolution text") {
		t.Error("output missing resolution text")
	}
}

func TestFormatResolucaoExerciciosInteger(t *testing.T) {
	buf, cleanup := captureStdout(t)
	FormatResolucaoExercicios(42)
	cleanup()

	output := buf.String()
	if !strings.Contains(output, "42") {
		t.Error("output missing integer value")
	}
}

func TestFormatResolucaoExerciciosMultipleLines(t *testing.T) {
	buf, cleanup := captureStdout(t)
	FormatResolucaoExercicios("line1\nline2\nline3")
	cleanup()

	output := buf.String()
	if !strings.Contains(output, "line1") {
		t.Error("output missing first line")
	}
	if !strings.Contains(output, "line2") {
		t.Error("output missing second line")
	}
	if !strings.Contains(output, "line3") {
		t.Error("output missing third line")
	}
}

func TestCreate(t *testing.T) {
	qs := QuestionnaireService{}
	input := []Questionnaire{
		{Number: "1.", Question: "Q1?", Options: "[a] [b]"},
		{Number: "2.", Question: "Q2?", Options: "[c] [d]"},
	}

	result := qs.Create(input)

	if len(result) != 2 {
		t.Fatalf("expected 2 questions, got %d", len(result))
	}
	if result[0].Number != "1." {
		t.Errorf("expected number '1.', got %q", result[0].Number)
	}
	if result[0].Question != "Q1?" {
		t.Errorf("expected question 'Q1?', got %q", result[0].Question)
	}
	if result[0].Options != "[a] [b]" {
		t.Errorf("expected options '[a] [b]', got %q", result[0].Options)
	}
}

func TestCreateEmpty(t *testing.T) {
	qs := QuestionnaireService{}
	result := qs.Create([]Questionnaire{})

	if len(result) != 0 {
		t.Errorf("expected empty result, got %d items", len(result))
	}
}
