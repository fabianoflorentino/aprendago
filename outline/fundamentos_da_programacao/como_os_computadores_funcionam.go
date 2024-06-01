package outline

import format "github.com/fabianoflorentino/aprendago/outline/format"

func ComoOsComputadoresFuncionam() {
	topic := format.OutlineContent{
		Title: "Como os computadores funcionam",
		Content: `
- Isso é importante pois daqui pra frente vamos falar de ints, bytes, e etc.
- Não é necessário um conhecimento a fundo mas é importante ter uma idéia de como as coisas funcionam por trás dos panos.
- https://docs.google.com/presentation/d/1aVytiGOBVDMISFW-ZARJ5iFY1osU2XJIw0hQpNICXm8/
- ASCII: https://en.wikipedia.org/wiki/ASCII
- Filme: Alan Turing, The Immitation Game.
    `,
	}

	format.FormatOutlineTopic(topic)
}
