package fundamentos_da_programacao

import (
	"testing"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

func TestChapterRegistration(t *testing.T) {
	c := chapter.Get(4)
	if c == nil {
		t.Errorf("expected chapter 4 to be registered")
	}
	if c.Number != 4 {
		t.Errorf("expected chapter number 4, got %d", c.Number)
	}
}

