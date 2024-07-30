package visao_geral_do_curso

import "github.com/fabianoflorentino/aprendago/pkg/format"

func Sucesso() {
	format.StartSection(func(outline *format.OutlineContent) {
		outline.AddHeader("Sucesso")
		outline.StartList(func(list *format.List) {
			list.AddItem("Qual o motivo que algumas pessoa obtem sucesso e outras não?")
			list.AddItem("Eu quero que você obtenha sucesso com este curso, então vamos falar sobre o assunto.")
			list.AddItem("Perguntaram a Bill Gates e Warren Buffet, independentemente, qual seria sua principal característica responsável pelo seu sucesso. A resposta de ambos foi:")
			list.StartSubList(func(list *format.List) {
				list.AddItem("Foco.")
				list.AddItem("(Fonte: https://www.linkedin.com/pulse/20140707144749-8353952-the-one-word-answer-to-why-bill-gates-and-warren-buffett-have-been-so-successful/)")
			})
			list.AddItem("O que isso quer dizer?")
			list.AddItem("Escolha um objetivo e se concentre nele. Faça desse objetivo sua maior prioridade.")
			list.AddItem(`Foco. Concentração. → "Vou conseguir chegar lá ou vou morrer tentando."`)
			list.AddItem("Exemplo: minha carreira de violinista.")
			list.AddItem(`Faça disso uma prioridade, não uma resolução de ano novo. Tem vezes que seus amigos vão te chamar pra fazer algo massa e você vai ter que dizer, "gente, não vai dar, estou estudando."`)
			list.AddItem("O Todd costuma dizer: de gota em gota se enche um balde. A cada dia uma gota. Algumas semanas depois você olha o balde, e pô meu não tem quase nada ali, pra que esse esforço todo? Mas depois de alguns anos o balde vai estar transbordando. É só trabalhar, de gota em gota.")
			list.AddItem("Então:")
			list.StartSubList(func(list *format.List) {
				list.AddItem("Seja proativo.")
				list.AddItem("Trabalhe. Invista as horas de esforço.")
				list.AddItem("Faça exercício, coma direito, tenha uma atitude positiva.")
				list.AddItem("Easy mode: pare de ver TV e use esse tempo pra estudar.")
			})
			list.AddItem("E algumas dicas práticas:")
			list.StartSubList(func(list *format.List) {
				list.AddItem("Metade das pessoas que assiste esse curso diz que eu falo muito rápido. A outra metade diz que eu falo muito devagar. Então use o botãozinho ali no player e selecione a velocidade mais apropriada pra você.")
				list.AddItem(`Boa parte de qualquer aprendizado se resume a memória muscular. Então quando eu passar exercícios, digite o código. Não copie e cole. Digite. E quando eu estiver programando "ao vivo," não passe pra frente, assista, e preste atenção. Se acostume com o processo, com o ato de programar.`)
			})
			list.AddItem("Com paciência e com persistência você chega lá.")
		})
	})
}
