package variaveis_valores_tipos

import (
	"testing"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

func TestChapterRegistration(t *testing.T) {
	c := chapter.Get(2)
	if c == nil {
		t.Errorf("expected chapter 2 to be registered")
	}
	if c.Number != 2 {
		t.Errorf("expected chapter number 2, got %d", c.Number)
	}
}

