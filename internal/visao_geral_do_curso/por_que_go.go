package visao_geral_do_curso

import "github.com/fabianoflorentino/aprendago/pkg/format"

func PorQueGo() {
	topic := format.OutlineContent{
		Title: "Por que Go?",
		Content: `
- Antes de investir seu tempo em aprender a linguagem Go, é bom você entender por que isso é uma boa idéia.
- O que estava acontecendo no Google...
- Criada por Ken Thompson (Unix, B, C), Rob Pike (UTF-8), e Robert Griesemer.
- Em 2006, não tinha uma linguagem de compilação rápida, execução rápida, e fácil de programar. É uma linguagem criada para resolver as questões de performance e complexidade.
- https://golang.org/doc/faq#creating_a_new_language
- Eficiente
  - Standard library é déis
  - Multiplataforma.
  - Garbage collection (lightning fast!)
  - Cross-compile.
- Fácil de usar
  - É uma linguagem compilada, de tipagem forte e estática,
  - Tem pouquíssimas palavras reservadas, que vamos aprender todas no curso, ou seja, é muito de boas de aprender
  - só sobe nas listas de popularidade
- Killer feature: Em 2006, logo após o primeiro dual core. Thread: 1mb. Goroutine: 2kb.
- É massa!
- Quando usar Go?
  - Escala
  - Seviços web, redes, servers (machine learning, image processing, crypto, ...)
  - Quando precisar de uma lingaugem rápida, simples, fácil de aprender, e fácil de usar.
- Usa em: APIs, CLIs, microservices, libraries/framework, processamento de dados, ... É a base dos serviços de cloud e orquestração de containers.
- Quem usa? https://github.com/golang/go/wiki/GoUsers
- Go é OOP? https://golang.org/doc/faq#Is_Go_an_object-oriented_language
- Mais um: https://golang.org/doc/faq#principles
- $$$: https://insights.stackoverflow.com/survey/2017#technology-top-paying-technologies-by-region -> US
`,
	}

	format.FormatOutlineTopic(topic)
}
