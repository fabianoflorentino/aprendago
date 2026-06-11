package teste_benchmarking

import (
	"testing"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

func TestChapterRegistration(t *testing.T) {
	c := chapter.Get(27)
	if c == nil {
		t.Errorf("expected chapter 27 to be registered")
	}
	if c.Number != 27 {
		t.Errorf("expected chapter number 27, got %d", c.Number)
	}
}

