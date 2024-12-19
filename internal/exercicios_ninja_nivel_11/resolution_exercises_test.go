package exercicios_ninja_nivel_11

import (
	"encoding/json"
	"testing"
)

func TestResolucaoNaPraticaExercicio1(t *testing.T) {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := `{"First":"James","Last":"Bond","Sayings":["Shaken, not stirred","Any last wishes?","Never say never"]}`
	if string(bs) != expected {
		t.Errorf("Expected %v, got %v", expected, string(bs))
	}
}

func TestResolucaoNaPraticaExercicio2(t *testing.T) {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := `{"First":"James","Last":"Bond","Sayings":["Shaken, not stirred","Any last wishes?","Never say never"]}`

	if string(bs) != expected {
		t.Errorf("Expected %v, got %v", expected, string(bs))
	}
}
