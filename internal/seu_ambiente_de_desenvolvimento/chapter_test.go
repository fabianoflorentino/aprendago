package seu_ambiente_de_desenvolvimento

import (
	"testing"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

func TestChapterRegistration(t *testing.T) {
	c := chapter.Get(19)
	if c == nil {
		t.Errorf("expected chapter 19 to be registered")
	}
	if c.Number != 19 {
		t.Errorf("expected chapter number 19, got %d", c.Number)
	}
}

