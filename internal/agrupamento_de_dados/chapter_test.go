package agrupamento_de_dados

import (
	"testing"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

func TestChapterRegistration(t *testing.T) {
	c := chapter.Get(8)
	if c == nil {
		t.Errorf("expected chapter 8 to be registered")
	}
	if c.Number != 8 {
		t.Errorf("expected chapter number 8, got %d", c.Number)
	}
}

