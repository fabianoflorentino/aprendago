package documentacao

import (
	"testing"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

func TestChapterRegistration(t *testing.T) {
	c := chapter.Get(25)
	if c == nil {
		t.Errorf("expected chapter 25 to be registered")
	}
	if c.Number != 25 {
		t.Errorf("expected chapter number 25, got %d", c.Number)
	}
}

