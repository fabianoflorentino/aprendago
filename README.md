# Aprenda Go

[Github](https://github.com/vkorbes/aprendago) [Youtube](https://youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&si=_JIbmByhwYvHdJAr)

## Docker

### Build

```shell
docker compose build --no-cache
```

### Run

```shell
docker compose up -d
```

### Stop

```shell
docker compose down
```

### Logs

```shell
docker compose logs -f
```

## Uso

### Container

```shell
docker compose exec -it aprendago /bin/sh -c 'go run main.go --help'
```

### Local

```shell
go run main.go --help
```

```shell
Uso: go run main.go [opção]

Exemplos:
  go run main.go --bem-vindo

Opções:
  --bem-vindo                                      Exibe a mensagem de boas-vindas ao curso Aprenda Go.
  --porque-go                                      Descreve os benefícios e razões para aprender a linguagem Go.
  --sucesso                                        Apresenta dicas e estratégias para ter sucesso no curso.
  --recursos                                       Lista recursos e materiais de apoio para o curso.
  --como-esse-curso-funciona                       Explica a estrutura e metodologia do curso.
  --go-playground                                  Fornece informações sobre o uso do Go Playground.
  --hello-world                                    Detalha o primeiro programa em Go: 'Hello, World!'.
  --operador-curto-de-declaracao                   Explica o uso do operador curto de declaração em Go.
  --a-palavra-reservada-var                        Descreve o uso da palavra reservada 'var' em Go.
  --explorando-tipos                               Explora os diferentes tipos de dados em Go.
  --valor-zero                                     Discute o conceito de valor zero em Go.
  --o-pacote-fmt                                   Fornece um resumo do pacote 'fmt' para formatação de E/S.
  --criando-seu-proprio-tipo                       Detalha como criar seus próprios tipos em Go.
  --conversao-nao-coercao                          Explica a diferença entre conversão e coerção em Go.
  --contribua-seu-codigo                           Fornece informações sobre como contribuir com seu próprio código.
  --na-pratica-exercicio-1 --nivel-1               Apresenta o primeiro exercício prático do curso.
  --na-pratica-exercicio-1 --nivel-1 --resolucao   Exibe a resolução do primeiro exercício prático.
  --na-pratica-exercicio-2 --nivel-1               Apresenta o segundo exercício prático do curso.
  --na-pratica-exercicio-2 --nivel-1 --resolucao   Exibe a resolução do segundo exercício prático.
  --na-pratica-exercicio-3 --nivel-1               Apresenta o terceiro exercício prático do curso.
  --na-pratica-exercicio-3 --nivel-1 --resolucao   Exibe a resolução do terceiro exercício prático.
  --na-pratica-exercicio-4 --nivel-1               Apresenta o quarto exercício prático do curso.
  --na-pratica-exercicio-4 --nivel-1 --resolucao   Exibe a resolução do quarto exercício prático.
  --na-pratica-exercicio-5 --nivel-1               Apresenta o quinto exercício prático do curso.
  --na-pratica-exercicio-5 --nivel-1 --resolucao   Exibe a resolução do quinto exercício prático.
  --na-pratica-exercicio-6 --nivel-1               Apresenta o sexto exercício prático do curso.
  --na-pratica-exercicio-6 --nivel-1 --prova       Exibe a prova do sexto exercício prático.
  --tipo-booleano                                  Explora o tipo de dados booleano em Go.
  --como-os-computadores-funcionam                 Descreve o funcionamento dos computadores e sua importância para a programação.
  --tipos-numericos                                Explora os tipos numéricos em Go.
  --overflow                                       Discute o conceito de overflow e como ele pode afetar seu código.
  --tipo-string                                    Explora o tipo de dados string em Go.
  --sistemas-numericos                             Apresenta os sistemas numéricos e sua importância para a programação.
  --constantes                                     Detalha o uso de constantes em Go.
  --iota                                           Explora o uso do identificador iota em Go.
  --deslocamento-de-bits                           Discute o conceito de deslocamento de bits em Go.
  --na-pratica-exercicio-1 --nivel-2               Apresenta o primeiro exercício prático do nível 2.
  --na-pratica-exercicio-1 --nivel-2 --resolucao   Exibe a resolução do primeiro exercício prático do nível 2.
  --outline                                        Exibe o outline completo do curso.
  --help                                           Exibe a lista de todas as opções disponíveis.
```

### Uso Contianer

```shell
docker compose exec -it aprendago /bin/sh -c 'go run main.go --bem-vindo'
```

### Uso Local

```shell
go run main.go --bem-vindo
```

```shell
Aprenda GO

Bem vindo

- Bem vindo ao curso!
- Eu sou...
- Go foi criado por gente foda que criou o Unix, B, UTF-8...
- Em 2006 o google queria...
- É uma lingguagem que vem crescrendo horrores...
- Nesse curso você vai aprender...
- O curriculo que vamos estudar...
- Para os novos na programação... Para os programadores experientes...
- Participe!
```
