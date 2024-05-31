package outline

import format "github.com/fabianoflorentino/aprendago/outline/format"

func Sucesso() {
	topic := format.OulineContent{
		Title: "Sucesso",
		Content: `
- Qual o motivo que algumas pessoa obtem sucesso e outras não?
- Eu quero que você obtenha sucesso com este curso, então vamos falar sobre o assunto.
- Perguntaram a Bill Gates e Warren Buffet, independentemente, qual seria sua principal característica responsável pelo seu sucesso. A resposta de ambos foi:
- Foco.
- (Fonte: https://www.linkedin.com/pulse/20140707144749-8353952-the-one-word-answer-to-why-bill-gates-and-warren-buffett-have-been-so-successful/)
- O que isso quer dizer?
- Escolha um objetivo e se concentre nele. Faça desse objetivo sua maior prioridade.
- Foco. Concentração. → "Vou conseguir chegar lá ou vou morrer tentando."
- Exemplo: minha carreira de violinista.
- Faça disso uma prioridade, não uma resolução de ano novo. Tem vezes que seus amigos vão te chamar pra fazer algo massa e você vai ter que dizer, "gente, não vai dar, estou estudando."
- O Todd costuma dizer: de gota em gota se enche um balde. A cada dia uma gota. Algumas semanas depois você olha o balde, e pô meu não tem quase nada ali, pra que esse esforço todo? Mas depois de alguns anos o balde vai estar transbordando. É só trabalhar, de gota em gota.

- Então:
  - Seja proativo.
  - Trabalhe. Invista as horas de esforço.
  - Faça exercício, coma direito, tenha uma atitude positiva.
  - Easy mode: pare de ver TV e use esse tempo pra estudar.

- E algumas dicas práticas:
  - Metade das pessoas que assiste esse curso diz que eu falo muito rápido. A outra metade diz que eu falo muito devagar. Então use o botãozinho ali no player e selecione a velocidade mais apropriada pra você.
  - Boa parte de qualquer aprendizado se resume a memória muscular. Então quando eu passar exercícios, digite o código. Não copie e cole. Digite. E quando eu estiver programando "ao vivo," não passe pra frente, assista, e preste atenção. Se acostume com o processo, com o ato de programar.

- Com paciência e com persistência você chega lá.
      `,
	}

	format.FormatOutlineTopic(topic)
}
