package aplicacoes

import (
	"encoding/json"
	"os"
)

type exampleJSON struct {
	Name     string
	Lastname string
	Age      int
}

func PrintExampleJSON() {
	toJSON, err := json.Marshal(showExampleJSON())
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(toJSON)
}

func showExampleJSON() exampleJSON {
	person := exampleJSON{
		Name:     "John",
		Lastname: "Doe",
		Age:      30,
	}

	return person
}
