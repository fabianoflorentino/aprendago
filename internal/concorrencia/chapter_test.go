package concorrencia

import (
	"testing"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

func TestChapterRegistration(t *testing.T) {
	c := chapter.Get(18)
	if c == nil {
		t.Errorf("expected chapter 18 to be registered")
	}
	if c.Number != 18 {
		t.Errorf("expected chapter number 18, got %d", c.Number)
	}
}

