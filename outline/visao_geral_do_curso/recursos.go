package visao_geral_do_curso

import format "github.com/fabianoflorentino/aprendago/outline/format"

func Recursos() {
	topic := format.OutlineContent{
		Title: "Recursos",
		Content: `
- Leia as descrições dos vídeos!
  - Outline completo: https://github.com/ellenkorbes/aprendago/blob/master/OUTLINE.md
- Exercícios:
- Comunidade:
  - Slack: https://invite.slack.golangbridge.org/
  - #brazil
  - #brazilian-go-studies, toda quinta às 22h!
    - gravações em: https://www.youtube.com/cesargimenes
    - exercícios em: https://github.com/crgimenes/Go-Hands-On
    - FB: https://www.facebook.com/gophers.br/
- Livros:
  - A Linguagem de Programação Go (Alan Donovan, Brian Kernighan): https://www.amazon.com.br/dp/8575225464/
  - Go em Ação (William Kennedy, Brian Ketelsen, Erik St. Martin): https://www.amazon.com.br/dp/8575225065/
  - Introdução à Linguagem Go (Caleb Doxsey): https://www.amazon.com.br/dp/8575224891/
- Em inglês: gobyexample.com, forum.golangbridge.org, stackoverflow.com/questions/tagged/go
- Documentação oficial:
  - golang.org
  - godoc.org
  - golang.org/doc/effective_go.html
- Podcast Go Time: https://changelog.com/gotime
- Defective Go: https://www.youtube.com/channel/UC98qIvCCqd4fjOw1ks78SwA
`,
	}

	format.FormatOutlineTopic(topic)
}
