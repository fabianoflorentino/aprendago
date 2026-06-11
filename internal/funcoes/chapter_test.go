package funcoes

import (
	"testing"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

func TestChapterRegistration(t *testing.T) {
	c := chapter.Get(12)
	if c == nil {
		t.Errorf("expected chapter 12 to be registered")
	}
	if c.Number != 12 {
		t.Errorf("expected chapter number 12, got %d", c.Number)
	}
}

func TestExamples(t *testing.T) {
	// Coverage: example functions are exercised via cmd integration tests
}
