package structs

import "github.com/fabianoflorentino/aprendago/outline/format"

func MenuStructs([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--structs", ExecFunc: func() { Struct() }},
		{Options: "--structs-embutidos", ExecFunc: func() { StructsEmbutidos() }},
		{Options: "--lendo-a-documentacao", ExecFunc: func() { LendoADocumentacao() }},
	}
}
