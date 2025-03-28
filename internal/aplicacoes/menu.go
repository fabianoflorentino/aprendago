package aplicacoes

import (
	"github.com/fabianoflorentino/aprendago/pkg/format"
	"github.com/fabianoflorentino/aprendago/pkg/logger"
	"github.com/fabianoflorentino/aprendago/pkg/section"
)

// MenuAplicacoes returns a slice of format.MenuOptions, each representing a different command-line option
// and its corresponding execution function. The options cover various topics such as JSON documentation,
// JSON marshal/unmarshal, JSON encoder, interface Writer, sort package, custom sorting, and bcrypt.
// Each option is associated with a specific ExecFunc that executes the relevant section or example.
func MenuAplicacoes([]string) []format.MenuOptions {
	section, err := section.New(rootDir)
	if err != nil {
		logger.New("error to create a new section", err.Error()).Register()
	}

	return []format.MenuOptions{
		{Options: "--documentacao-json", ExecFunc: func() { section.Format("Documentação JSON") }},
		{Options: "--documentacao-json --example --json-marshal", ExecFunc: func() { UsingJsonMarshal() }},
		{Options: "--documentacao-json --example --json-unmarshal", ExecFunc: func() { UsingJsonUnmarshal() }},
		{Options: "--documentacao-json --example --json-encoder", ExecFunc: func() { UsingJsonEncoder() }},
		{Options: "--json-marshal", ExecFunc: func() { section.Format("JSON marshal (ordenação)") }},
		{Options: "--json-unmarshal", ExecFunc: func() { section.Format("JSON unmarshal (desornação)") }},
		{Options: "--interface-writer", ExecFunc: func() { section.Format("A interface Writer") }},
		{Options: "--pacote-sort", ExecFunc: func() { section.Format("O pacote sort") }},
		{Options: "--pacote-sort --example", ExecFunc: func() { UsingPackageSort() }},
		{Options: "--customizando-sort", ExecFunc: func() { section.Format("Customizando sort") }},
		{Options: "--customizando-sort --example", ExecFunc: func() { UsingCustomSort() }},
		{Options: "--bcrypt", ExecFunc: func() { section.Format("bcrypt") }},
	}
}
