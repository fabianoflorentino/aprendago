package fundamentos_da_programacao

import "github.com/fabianoflorentino/aprendago/pkg/format"

func SistemasNumericos() {
	topic := format.OutlineContent{
		Title: "Sistemas Numéricos",
		Content: `
- Base-10: decimal, 0–9
- Base-2: binário, 0–1
- Base-16: hexadecimal, 0–f
- https://docs.google.com/document/d/1GqXpubhMMIr4Sy5xwgiPIDh5PGVmVpF2u0c9vDrvykE/
- Demonstração em Go.
    `,
	}

	format.FormatOutlineTopic(topic)
}
