package visao_geral_do_curso

import (
	"testing"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

func TestChapterRegistration(t *testing.T) {
	c := chapter.Get(1)
	if c == nil {
		t.Errorf("expected chapter 1 to be registered")
	}
	if c.Number != 1 {
		t.Errorf("expected chapter number 1, got %d", c.Number)
	}
}

