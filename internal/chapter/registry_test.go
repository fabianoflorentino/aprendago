package chapter

import (
	"testing"
)

func TestRegisterAndAll(t *testing.T) {
	// Reset registry for test isolation
	old := registry
	registry = nil
	defer func() { registry = old }()

	c1 := &Chapter{Number: 1, Title: "Cap 1", RootDir: "testdata"}
	c2 := &Chapter{Number: 2, Title: "Cap 2", RootDir: "testdata"}

	Register(c1)
	Register(c2)

	all := All()
	if len(all) != 2 {
		t.Fatalf("expected 2 registered chapters, got %d", len(all))
	}
	if all[0] != c1 {
		t.Error("expected first registered chapter to be c1")
	}
	if all[1] != c2 {
		t.Error("expected second registered chapter to be c2")
	}
}

func TestGet(t *testing.T) {
	// Reset registry for test isolation
	old := registry
	registry = nil
	defer func() { registry = old }()

	c1 := &Chapter{Number: 1, Title: "Cap 1", RootDir: "testdata"}
	c5 := &Chapter{Number: 5, Title: "Cap 5", RootDir: "testdata"}

	Register(c1)
	Register(c5)

	tests := []struct {
		number  int
		want    *Chapter
		wantNil bool
	}{
		{1, c1, false},
		{5, c5, false},
		{3, nil, true},
		{99, nil, true},
	}

	for _, tt := range tests {
		got := Get(tt.number)
		if tt.wantNil {
			if got != nil {
				t.Errorf("Get(%d) = %v, want nil", tt.number, got)
			}
		} else if got != tt.want {
			t.Errorf("Get(%d) = %v, want %v", tt.number, got, tt.want)
		}
	}
}

func TestRegisterEmptyAll(t *testing.T) {
	// Reset registry for test isolation
	old := registry
	registry = nil
	defer func() { registry = old }()

	all := All()
	if len(all) != 0 {
		t.Errorf("expected empty registry, got %d chapters", len(all))
	}

	if got := Get(1); got != nil {
		t.Errorf("expected nil for Get(1) on empty registry, got %v", got)
	}
}
