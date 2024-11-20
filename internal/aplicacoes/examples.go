package aplicacoes

import (
	"encoding/json"
	"fmt"
	"os"
)

type pessoa struct {
	Name     string
	Lastname string
	Age      int
}

type extractInfo struct {
	Name     string `json:"Name"`
	Lastname string `json:"Lastname"`
	Age      int    `json:"Age"`
}

func UsingJsonMarshal() {
	p := pessoa{
		Name:     "Fabiano",
		Lastname: "Florentino",
		Age:      40,
	}

	pessoaToJSON, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(pessoaToJSON)
}

func UsingJsonUnmarshal() {
	pessoaJSON := []byte(`{"Name":"Fabiano","Lastname":"Florentino","Age":40}`)

	var p extractInfo
	err := json.Unmarshal(pessoaJSON, &p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Name: %s\nLastname: %s\nAge: %d\n", p.Name, p.Lastname, p.Age)
}

func UsingJsonEncoder() {
	p := pessoa{
		Name:     "Fabiano",
		Lastname: "Florentino",
		Age:      40,
	}

	json.NewEncoder(os.Stdout).Encode(p)
}
