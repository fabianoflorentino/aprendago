package structs

import (
	"testing"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

func TestChapterRegistration(t *testing.T) {
	c := chapter.Get(10)
	if c == nil {
		t.Errorf("expected chapter 10 to be registered")
	}
	if c.Number != 10 {
		t.Errorf("expected chapter number 10, got %d", c.Number)
	}
}

