package fluxo_de_controle

import (
	"testing"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

func TestChapterRegistration(t *testing.T) {
	c := chapter.Get(6)
	if c == nil {
		t.Errorf("expected chapter 6 to be registered")
	}
	if c.Number != 6 {
		t.Errorf("expected chapter number 6, got %d", c.Number)
	}
}

