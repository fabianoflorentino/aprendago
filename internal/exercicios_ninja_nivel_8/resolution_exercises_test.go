package exercicios_ninja_nivel_8

import (
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"github.com/fabianoflorentino/aprendago/pkg/output"
	"github.com/fabianoflorentino/aprendago/pkg/trim"
)

const (
	expectTemplate = "\nwant:\n%s\n\ngot:\n%s\n"
)

func TestResolucaoNaPraticaExercicio1(t *testing.T) {
	output := output.New()
	result, err := output.Capture(ResolucaoNaPraticaExercicio1)
	if err != nil {
		logger.Log("Failed to capture output: %v", err)
	}

	expect := `
Exemplo 1

User 1: {"First":"M","Age":54}
User 2: {"First":"Q","Age":64}
User 3: {"First":"Moneypenny","Age":27}
User 4: {"First":"James","Age":32}

Exemplo 2

[{"First":"M","Age":54},{"First":"Q","Age":64},{"First":"Moneypenny","Age":27},{"First":"James","Age":32}]
`

	trim := trim.New()

	if !strings.Contains(trim.String(result), trim.String(expect)) {
		t.Errorf(expectTemplate, expect, result)
	}
}
