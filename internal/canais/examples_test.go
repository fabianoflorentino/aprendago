package canais

import (
	"strings"
	"testing"

	"github.com/fabianoflorentino/aprendago/pkg/output"
)

func TestUsingConverge(t *testing.T) {
	out := output.New()
	result, err := out.Capture(UsingConverge)
	if err != nil {
		t.Fatalf("failed to capture output: %v", err)
	}

	if !strings.Contains(result, "Even Value: 0") {
		t.Errorf("expected output to contain 'Even Value: 0', got: %s", result)
	}

	if !strings.Contains(result, "Odd Value: 99") {
		t.Errorf("expected output to contain 'Odd Value: 99', got: %s", result)
	}
}

func TestUsingDivergence(t *testing.T) {
	out := output.New()
	result, err := out.Capture(UsingDivergence)
	if err != nil {
		t.Fatalf("failed to capture output: %v", err)
	}

	if !strings.Contains(result, "Received: 0") {
		t.Errorf("expected output to contain 'Received: 0', got: %s", result)
	}
}

func TestSendToDivergence(t *testing.T) {
	chann := make(chan int)
	go sendToDivergence(5, chann)

	var received []int
	for value := range chann {
		received = append(received, value)
	}

	if len(received) != 5 {
		t.Errorf("expected 5 values, got %d", len(received))
	}

	for i := 0; i < 5; i++ {
		if received[i] != i {
			t.Errorf("expected received[%d] = %d, got %d", i, i, received[i])
		}
	}
}
