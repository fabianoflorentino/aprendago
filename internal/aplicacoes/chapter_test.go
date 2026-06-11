package aplicacoes

import (
	"testing"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

func TestChapterRegistration(t *testing.T) {
	c := chapter.Get(16)
	if c == nil {
		t.Errorf("expected chapter 16 to be registered")
	}
	if c.Number != 16 {
		t.Errorf("expected chapter number 16, got %d", c.Number)
	}
}

func TestExamples(t *testing.T) {
	// Coverage: example functions are exercised via cmd integration tests
}
