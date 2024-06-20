package structs

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/outline/format"
)

func HelpMeStructs() {
	hlp := []format.HelpMe{
		{Flag: "--structs", Description: "Structs", Width: 0},
		{Flag: "--structs-embutidos", Description: "Structs Embutidos", Width: 0},
		{Flag: "--lendo-a-documentacao", Description: "Lendo a documentação", Width: 0},
	}

	fmt.Println("\n Capítulo 10: Structs")
	format.PrintHelpMe(hlp)
}
