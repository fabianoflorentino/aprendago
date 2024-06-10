package agrupamento_de_dados

import format "github.com/fabianoflorentino/aprendago/outline/format"

func Array() {
	topic := format.OutlineContent{
		Title: "Array",
		Content: `
- Estruturas de dados, ou agrupamentos de dados, nos permitem agrupar valores diferentes. Estes valores podem ser ou não do mesmo tipo.
- As estruturas que veremos são: arrays, slices, structs e maps.
- Vamos começar com arrays. Arrays em Go são uma fundação, e não algo que utilizamos todo dia.
- Seu tamanho deve estar presente na declaração: var x [n]int
- Atribui-se valores a suas posições com: x[i] = y (0-based)
- Para ver o tamanho usa-se: len(x)
- ref/spec: "The length is part of the array's type" → [5]int != [6]int
- Effective Go: Arrays são úteis para [umas coisas que a gente não vai fazer nunca] e servem de fundação para slices. Use slices ao invés de arrays.
- Go Playground: https://play.golang.org/p/Fv-sDF-ryZ
    `,
	}

	format.FormatOutlineTopic(topic)
}
