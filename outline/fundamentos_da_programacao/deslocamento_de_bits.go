package outline

import format "github.com/fabianoflorentino/aprendago/outline/format"

func DeslocamentoDeBits() {
	topic := format.OutlineContent{
		Title: "Deslocamento de Bits",
		Content: `
- Deslocamento de bits é quando deslocamos digitos binários para a esquerda ou direita.
- Na prática:
  - %d %b
  - x << y
  - iota * 10 << 10 = kb, mb, gb

- https://play.golang.org/p/7MOnbhx4R4
- https://splice.com/blog/iota-elegant-constants-golang/
- https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827

- Fim da sessão. Massa!
  `,
	}

	format.FormatOutlineTopic(topic)
}
