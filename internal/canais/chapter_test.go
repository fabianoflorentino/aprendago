package canais

import (
	"testing"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

func TestChapterRegistration(t *testing.T) {
	c := chapter.Get(21)
	if c == nil {
		t.Errorf("expected chapter 21 to be registered")
	}
	if c.Number != 21 {
		t.Errorf("expected chapter number 21, got %d", c.Number)
	}
}

func TestExamples(t *testing.T) {
	// Coverage: example functions are exercised via cmd integration tests
}
