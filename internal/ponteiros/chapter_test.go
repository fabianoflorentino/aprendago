package ponteiros

import (
	"testing"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

func TestChapterRegistration(t *testing.T) {
	c := chapter.Get(14)
	if c == nil {
		t.Errorf("expected chapter 14 to be registered")
	}
	if c.Number != 14 {
		t.Errorf("expected chapter number 14, got %d", c.Number)
	}
}

