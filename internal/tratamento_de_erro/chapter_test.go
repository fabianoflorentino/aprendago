package tratamento_de_erro

import (
	"testing"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

func TestChapterRegistration(t *testing.T) {
	c := chapter.Get(23)
	if c == nil {
		t.Errorf("expected chapter 23 to be registered")
	}
	if c.Number != 23 {
		t.Errorf("expected chapter number 23, got %d", c.Number)
	}
}

