# Aprenda Go

[![Build, Publish, Tag and Release](https://github.com/fabianoflorentino/aprendago/actions/workflows/ci.yml/badge.svg)](https://github.com/fabianoflorentino/aprendago/actions/workflows/ci.yml) [![CodeQL](https://github.com/fabianoflorentino/aprendago/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/fabianoflorentino/aprendago/actions/workflows/github-code-scanning/codeql)

## Origem do Projeto

Curso de Go para iniciantes. Aprenda Go é um curso abrangente que ensina desde o básico até tópicos avançados da linguagem Go. Ministrado por [Ellen Körbes](https://www.linkedin.com/in/vkorbes/), uma entusiasta da linguagem, o curso é gratuito e está disponível no YouTube e no GitHub. Acesse a playlist no YouTube [aqui](https://youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&si=_JIbmByhwYvHdJAr) e o repositório no GitHub [aqui](https://github.com/vkorbes/aprendago).

## Objetivo

Este projeto tem como objetivo criar um CLI para facilitar a navegação e o acesso ao conteúdo do curso Aprenda Go. O CLI foi desenvolvido em Go e utiliza as bibliotecas padrão da linguagem para criar um menu interativo que permite ao usuário acessar os tópicos do curso. O CLI esta sendo desenvolvido como parte da prática de estudos da linguagem Go.

## Afim de contribuir?

Se você deseja contribuir com o projeto, fique à vontade para abrir uma [issue](https://github.com/fabianoflorentino/aprendago/issues) ou submeter um [pull request](https://github.com/fabianoflorentino/aprendago/pulls). Sua contribuição é muito bem-vinda!

## Montando o Ambiente de Desenvolvimento

Faça o clone do repositório:

```shell
git clone https://github.com/fabianoflorentino/aprendago.git
```

### Construa a imagem Docker

```shell
docker compose build --no-cache
```

### Inicie o container

```shell
docker compose up -d
```

### Parando o container

```shell
docker compose down
```

## Testando o CLI

### Menu Caps (Capítulos)

```shell
docker compose exec -it aprendago /bin/sh -c 'go run cmd/aprendago/main.go --caps'
```

```shell
Aprenda GO

Capítulos do Curso

  --cap-1 --topics    Visão Geral do Curso
  --cap-2 --topics    Variáveis, Valores e Tipos
  --cap-3 --topics    Exercícios Ninja: Nível 1
  --cap-4 --topics    Fundamentos da Programação
  --cap-5 --topics    Exercícios Ninja: Nível 2
  --cap-6 --topics    Fluxo de Controle
  --cap-7 --topics    Exercícios Ninja: Nível 3
  --cap-8 --topics    Agrupamento de Dados
  --cap-9 --topics    Exercícios Ninja: Nível 4
  --cap-10 --topics   Structs
  --cap-11 --topics   Exercícios Ninja: Nível 5
  --cap-12 --topics   Funções
  --cap-13 --topics   Exercícios Ninja: Nível 6
```

### Ajuda

```shell
docker compose exec -it aprendago /bin/sh -c 'go run cmd/aprendago/main.go --help'
```

```shell
Aprenda GO


Uso: go run main.go [opção]

Exemplo:
  go run main.go --bem-vindo

Ajuda:

--outline  Exibe o outline completo do curso.
--help     Exibe a lista de todas as opções disponíveis.
--caps     Exibe a lista de capítulos disponíveis.

Capítulos do Curso

  --cap-1 --topics    Visão Geral do Curso
  --cap-2 --topics    Variáveis, Valores e Tipos
  --cap-3 --topics    Exercícios Ninja: Nível 1
  --cap-4 --topics    Fundamentos da Programação
  --cap-5 --topics    Exercícios Ninja: Nível 2
  --cap-6 --topics    Fluxo de Controle
  --cap-7 --topics    Exercícios Ninja: Nível 3
  --cap-8 --topics    Agrupamento de Dados
  --cap-9 --topics    Exercícios Ninja: Nível 4
  --cap-10 --topics   Structs
  --cap-11 --topics   Exercícios Ninja: Nível 5
  --cap-12 --topics   Funções
  --cap-13 --topics   Exercícios Ninja: Nível 6

Outline do Curso por Capítulo

  --cap-1 --overview   Visão Geral do Curso
  --cap-2 --overview   Variáveis, Valores &amp; Tipos
  --cap-3 --overview   Exercícios Ninja Nível 1
  --cap-4 --overview   Fundamentos da Programação
  --cap-5 --overview   Exercícios Ninja Nível 2
  --cap-6 --overview   Fluxo de Controle

Capítulo 1: Visão Geral do Curso

  --bem-vindo                  Exibe a mensagem de boas-vindas ao curso Aprenda Go.
  --porque-go                  Descreve os benefícios e razões para aprender a linguagem Go.
  --sucesso                    Apresenta dicas e estratégias para ter sucesso no curso.
  --recursos                   Lista recursos e materiais de apoio para o curso.
  --como-esse-curso-funciona   Explica a estrutura e metodologia do curso.

Capítulo 2: Variáveis, Valores e Tipos

  --go-playground                  Exibe o Go Playground.
  --hello-world                    Exibe o Hello World.
  --operador-curto-de-declaracao   Exibe o operador curto de declaração.
  --a-palavra-reservada-var        Exibe a palavra reservada var.
  --explorando-tipos               Exibe como explorar tipos.
  --valor-zero                     Exibe o valor zero.
  --o-pacote-fmt                   Exibe o pacote fmt.
  --criando-seu-proprio-tipo       Exibe como criar seu próprio tipo.
  --conversao-nao-coercao          Exibe a conversão não coerção.

Capítulo 3: Exercícios Ninja Nível 1

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

Capítulo 4: Fundamentos da Programação

  --tipo-booleano                    Explora o tipo de dados booleano em Go.
  --como-os-computadores-funcionam   Descreve o funcionamento dos computadores e sua importância para a programação.
  --tipos-numericos                  Explora os tipos numéricos em Go.
  --overflow                         Discute o conceito de overflow e como ele pode afetar seu código.
  --tipo-string                      Explora o tipo de dados string em Go.
  --sistemas-numericos               Apresenta os sistemas numéricos e sua importância para a programação.
  --constantes                       Detalha o uso de constantes em Go.
  --iota                             Explora o uso do identificador iota em Go.
  --deslocamento-de-bits             Discute o conceito de deslocamento de bits em Go.

Capítulo 5: Exercícios Ninja - Nível 2

  --na-pratica-exercicio-1 --nivel-2               Apresenta o primeiro exercício prático do nível 2.
  --na-pratica-exercicio-1 --nivel-2 --resolucao   Exibe a resolução do primeiro exercício prático do nível 2.
  --na-pratica-exercicio-2 --nivel-2               Apresenta o segundo exercício prático do nível 2.
  --na-pratica-exercicio-2 --nivel-2 --resolucao   Exibe a resolução do segundo exercício prático do nível 2.
  --na-pratica-exercicio-3 --nivel-2               Apresenta o terceiro exercício prático do nível 2.
  --na-pratica-exercicio-3 --nivel-2 --resolucao   Exibe a resolução do terceiro exercício prático do nível 2.
  --na-pratica-exercicio-4 --nivel-2               Apresenta o quarto exercício prático do nível 2.
  --na-pratica-exercicio-4 --nivel-2 --resolucao   Exibe a resolução do quarto exercício prático do nível 2.
  --na-pratica-exercicio-5 --nivel-2               Apresenta o quinto exercício prático do nível 2.
  --na-pratica-exercicio-5 --nivel-2 --resolucao   Exibe a resolução do quinto exercício prático do nível 2.
  --na-pratica-exercicio-6 --nivel-2               Apresenta o sexto exercício prático do nível 2.
  --na-pratica-exercicio-6 --nivel-2 --resolucao   Exibe a resolução do sexto exercício prático do nível 2.
  --na-pratica-exercicio-7 --nivel-2               Apresenta o sétimo exercício prático do nível 2.
  --na-pratica-exercicio-7 --nivel-2 --prova       Exibe a prova do sétimo exercício prático do nível 2.

Capípulo 6: Fluxo de Controle

  --entendendo-fluxo-de-controle           Explica o conceito de fluxo de controle em Go.
  --loops-inicializacao-condicao-pos       Detalha o uso de loops com inicialização, condição e pós em Go.
  --loops-nested-loop                      Explora o conceito de loops aninhados em Go.
  --loops-a-declaracao-for                 Apresenta a declaração for em Go.
  --loops-break-e-continue                 Discute as instruções break e continue em loops em Go.
  --loops-utilizando-ascii                 Desafio surpresa: utilize ASCII para exibir texto em Go.
  --loops-utilizando-ascii --resolucao     Desafio surpresa: utilize ASCII para exibir texto em Go.
  --condicionais-a-declaracao-if           Apresenta a declaração if em Go.
  --condicionais-if-else-if-else           Detalha a declaração if-else-if-else em Go.
  --condicionais-a-declaracao-switch       Apresenta a declaração switch em Go.
  --condicionais-a-declaracao-switch-pt2   Detalha a declaração switch em Go.
  --operadores-logicos-condicionais        Explora os operadores lógicos condicionais em Go.

Capítulo 7: Exercícios Ninja Nível 3

  --na-pratica-exercicio-1 --nivel-3                Apresenta o primeiro exercício prático do nível 3.
  --na-pratica-exercicio-1 --nivel-3 --resolucao    Exibe a resolução do primeiro exercício prático do nível 3.
  --na-pratica-exercicio-2 --nivel-3                Apresenta o segundo exercício prático do nível 3.
  --na-pratica-exercicio-2 --nivel-3 --resolucao    Exibe a resolução do segundo exercício prático do nível 3.
  --na-pratica-exercicio-3 --nivel-3                Apresenta o terceiro exercício prático do nível 3.
  --na-pratica-exercicio-3 --nivel-3 --resolucao    Exibe a resolução do terceiro exercício prático do nível 3.
  --na-pratica-exercicio-4 --nivel-3                Apresenta o quarto exercício prático do nível 3.
  --na-pratica-exercicio-4 --nivel-3 --resolucao    Exibe a resolução do quarto exercício prático do nível 3.
  --na-pratica-exercicio-5 --nivel-3                Apresenta o quinto exercício prático do nível 3.
  --na-pratica-exercicio-5 --nivel-3 --resolucao    Exibe a resolução do quinto exercício prático do nível 3.
  --na-pratica-exercicio-6 --nivel-3                Apresenta o sexto exercício prático do nível 3.
  --na-pratica-exercicio-6 --nivel-3 --resolucao    Exibe a resolução do sexto exercício prático do nível 3.
  --na-pratica-exercicio-7 --nivel-3                Apresenta o sétimo exercício prático do nível 3.
  --na-pratica-exercicio-7 --nivel-3 --resolucao    Exibe a resolução do sétimo exercício prático do nível 3.
  --na-pratica-exercicio-8 --nivel-3                Apresenta o oitavo exercício prático do nível 3.
  --na-pratica-exercicio-8 --nivel-3 --resolucao    Exibe a resolução do oitavo exercício prático do nível 3.
  --na-pratica-exercicio-9 --nivel-3                Apresenta o nono exercício prático do nível 3.
  --na-pratica-exercicio-9 --nivel-3 --resolucao    Exibe a resolução do nono exercício prático do nível 3.
  --na-pratica-exercicio-10 --nivel-3               Apresenta o décimo exercício prático do nível 3.
  --na-pratica-exercicio-10 --nivel-3 --resolucao   Exibe a resolução do décimo exercício prático do nível 3.

Capítulo 8: Agrupamento de Dados

  --array                                                  Apresenta o tópico Array.
  --slice-literal-composta                                 Apresenta o tópico Slice Literal Composta.
  --slice-for-range                                        Apresenta o tópico Slice: for range.
  --slice-fatiando-ou-deletando-de-uma-fatia               Apresenta o tópico Slice: fatiando ou deletando de uma fatia.
  --slice-fatiando-ou-deletando-de-uma-fatia --resolucao   Apresenta a resolução do tópico Slice: fatiando ou deletando de uma fatia.
  --slice-anexando-a-uma-slice                             Apresenta o tópico Slice: anexando a uma slice.
  --slice-make                                             Apresenta o tópico Slice: Make.
  --slice-multi-dimensional                                Apresenta o tópico Slice: Multi Dimensional.
  --slice-a-surpresa-do-array-subjacente                   Apresenta o tópico Slice: a surpresa do array subjacente.
  --maps-introducao                                        Apresenta o tópico Maps: introdução.
  --maps-range-e-deletando                                 Apresenta o tópico Maps: Range e Deletando.

Capítulo 9: Exercícios Ninja Nível 4

  --na-pratica-exercicio-1 --nivel-4                Apresenta o primeiro exercício prático do Nível 4.
  --na-pratica-exercicio-1 --nivel-4 --resolucao    Exibe a resolução do primeiro exercício prático do Nível 4.
  --na-pratica-exercicio-2 --nivel-4                Apresenta o segundo exercício prático do Nível 4.
  --na-pratica-exercicio-2 --nivel-4 --resolucao    Exibe a resolução do segundo exercício prático do Nível 4.
  --na-pratica-exercicio-3 --nivel-4                Apresenta o terceiro exercício prático do Nível 4.
  --na-pratica-exercicio-3 --nivel-4 --resolucao    Exibe a resolução do terceiro exercício prático do Nível 4.
  --na-pratica-exercicio-4 --nivel-4                Apresenta o quarto exercício prático do Nível 4.
  --na-pratica-exercicio-4 --nivel-4 --resolucao    Exibe a resolução do quarto exercício prático do Nível 4.
  --na-pratica-exercicio-5 --nivel-4                Apresenta o quinto exercício prático do Nível 4.
  --na-pratica-exercicio-5 --nivel-4 --resolucao    Exibe a resolução do quinto exercício prático do Nível 4.
  --na-pratica-exercicio-6 --nivel-4                Apresenta o sexto exercício prático do Nível 4.
  --na-pratica-exercicio-6 --nivel-4 --resolucao    Exibe a resolução do sexto exercício prático do Nível 4.
  --na-pratica-exercicio-7 --nivel-4                Apresenta o sétimo exercício prático do Nível 4.
  --na-pratica-exercicio-7 --nivel-4 --resolucao    Exibe a resolução do sétimo exercício prático do Nível 4.
  --na-pratica-exercicio-8 --nivel-4                Apresenta o oitavo exercício prático do Nível 4.
  --na-pratica-exercicio-8 --nivel-4 --resolucao    Exibe a resolução do oitavo exercício prático do Nível 4.
  --na-pratica-exercicio-9 --nivel-4                Apresenta o nono exercício prático do Nível 4.
  --na-pratica-exercicio-9 --nivel-4 --resolucao    Exibe a resolução do nono exercício prático do Nível 4.
  --na-pratica-exercicio-10 --nivel-4               Apresenta o décimo exercício prático do Nível 4.
  --na-pratica-exercicio-10 --nivel-4 --resolucao   Exibe a resolução do décimo exercício prático do Nível 4.

 Capítulo 10: Structs

  --structs                Structs
  --structs-embutidos      Structs Embutidos
  --lendo-a-documentacao   Lendo a documentação
  --structs-anonimos       Structs Anônimos

Capítulo 11: Exercícios Ninja Nível 5

  --na-pratica-exercicio-1 --nivel-5               Apresenta o primeiro exercício prático do Nível 5.
  --na-pratica-exercicio-1 --nivel-5 --resolucao   Exibe a resolução do primeiro exercício prático do Nível 5.
  --na-pratica-exercicio-2 --nivel-5               Apresenta o segundo exercício prático do Nível 5.
  --na-pratica-exercicio-2 --nivel-5 --resolucao   Exibe a resolução do segundo exercício prático do Nível 5.
  --na-pratica-exercicio-3 --nivel-5               Apresenta o terceiro exercício prático do Nível 5.
  --na-pratica-exercicio-3 --nivel-5 --resolucao   Exibe a resolução do terceiro exercício prático do Nível 5.
  --na-pratica-exercicio-4 --nivel-5               Apresenta o quarto exercício prático do Nível 5.
  --na-pratica-exercicio-4 --nivel-5 --resolucao   Exibe a resolução do quarto exercício prático do Nível 5.

Capítulo 12: Funções

  --sintaxe                            Sintaxe de declaração de função
  --desenroland-enumerando-uma-slice   Descreve como iterar (enumerar) uma slice
  --defer                              Descreve o uso da declaração defer
  --metodos                            Descreve o que são métodos em Go
  --interfaces-e-polimorfismo          Descreve o que são interfaces e polimorfismo em Go
  --funcoes-anonimas                   Descreve o que são funções anônimas em Go
  --func-como-expressa                 Descreve como declarar funções como expressões
  --retornando-uma-funcao              Descreve como retornar uma função em Go
  --callback                           Descreve o que é um callback em Go
  --resolucao-desafio-callback         Mostra a resolução do desafio de callback
  --closure                            Descreve o que é um closure em Go
  --recursividade                      Descreve o que é recursividade em Go

Capítulo 13: Exercícios Ninja Nível 6

  --na-pratica-exercicio-1 --nivel-6               Apresenta o primeiro exercício prático do Nível 6.
  --na-pratica-exercicio-1 --nivel-6 --resolucao   Apresenta a resolução do primeiro exercício prático do Nível 6.
  --na-pratica-exercicio-2 --nivel-6               Apresenta o segundo exercício prático do Nível 6.
  --na-pratica-exercicio-2 --nivel-6 --resolucao   Apresenta a resolução do segundo exercício prático do Nível 6.
  --na-pratica-exercicio-3 --nivel-6               Apresenta o terceiro exercício prático do Nível 6.
  --na-pratica-exercicio-3 --nivel-6 --resolucao   Apresenta a resolução do terceiro exercício prático do Nível 6.
  --na-pratica-exercicio-4 --nivel-6               Apresenta o quarto exercício prático do Nível 6.
  --na-pratica-exercicio-4 --nivel-6 --resolucao   Apresenta a resolução do quarto exercício prático do Nível 6.
  --na-pratica-exercicio-5 --nivel-6               Apresenta o quinto exercício prático do Nível 6.
  --na-pratica-exercicio-5 --nivel-6 --resolucao   Apresenta a resolução do quinto exercício prático do Nível 6.
```

## Estrutura do Projeto

```shell
.
├── LICENSE
├── README.md
├── build
│   ├── Dockerfile.dev
│   └── Dockerfile.prod
├── cmd
│   └── aprendago
│       └── main.go
├── configs
│   └── _env
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── agrupamento_de_dados
│   │   ├── overview.yml
│   │   └── topics.go
│   ├── exercicios_ninja_nivel_1
│   │   ├── overview.yml
│   │   ├── resolution_exercises.go
│   │   └── topics.go
│   ├── exercicios_ninja_nivel_2
│   │   ├── overview.yml
│   │   ├── resolution_exercises.go
│   │   └── topics.go
│   ├── exercicios_ninja_nivel_3
│   │   ├── overview.yml
│   │   ├── resolucao_exercicios.go
│   │   └── topics.go
│   ├── exercicios_ninja_nivel_4
│   │   ├── helpme_exercicios_ninja_nivel_4.go
│   │   ├── menu_exercicios_ninja_nivel_4.go
│   │   ├── na_pratica_exercicio_1.go
│   │   ├── na_pratica_exercicio_10.go
│   │   ├── na_pratica_exercicio_2.go
│   │   ├── na_pratica_exercicio_3.go
│   │   ├── na_pratica_exercicio_4.go
│   │   ├── na_pratica_exercicio_5.go
│   │   ├── na_pratica_exercicio_6.go
│   │   ├── na_pratica_exercicio_7.go
│   │   ├── na_pratica_exercicio_8.go
│   │   ├── na_pratica_exercicio_9.go
│   │   └── topics.go
│   ├── exercicios_ninja_nivel_5
│   │   ├── helpme_exercicios_ninja_nivel_5.go
│   │   ├── menu_exercicios_ninja_nivel_5.go
│   │   ├── na_pratica_exercicio_1.go
│   │   ├── na_pratica_exercicio_2.go
│   │   ├── na_pratica_exercicio_3.go
│   │   ├── na_pratica_exercicio_4.go
│   │   └── topics.go
│   ├── exercicios_ninja_nivel_6
│   │   ├── helpme_exercicicios_ninja_nivel_6.go
│   │   ├── menu_exercicios_ninja_nivel_6.go
│   │   ├── na_pratica_exercicio_1.go
│   │   ├── na_pratica_exercicio_2.go
│   │   ├── na_pratica_exercicio_3.go
│   │   ├── na_pratica_exercicio_4.go
│   │   ├── na_pratica_exercicio_5.go
│   │   └── topics.go
│   ├── fluxo_de_controle
│   │   ├── overview.yml
│   │   └── topics.go
│   ├── funcoes
│   │   ├── callback.go
│   │   ├── closure.go
│   │   ├── defer.go
│   │   ├── desenrolando_enumerando_uma_slice.go
│   │   ├── func_como_expressao.go
│   │   ├── funcoes_anonimas.go
│   │   ├── helpme_funcoes.go
│   │   ├── interfaces_e_polimorfismo.go
│   │   ├── menu_funcoes.go
│   │   ├── metodos.go
│   │   ├── recursividade.go
│   │   ├── retornando_uma_funcao.go
│   │   ├── sintaxe.go
│   │   └── topics.go
│   ├── fundamentos_da_programacao
│   │   ├── overview.yml
│   │   └── topics.go
│   ├── menu
│   │   ├── capitulo_options.go
│   │   ├── capitulo_outline.go
│   │   ├── helpme.go
│   │   └── options.go
│   ├── outline.go
│   ├── structs
│   │   ├── overview.yml
│   │   └── topics.go
│   ├── variaveis_valores_tipos
│   │   ├── overview.yml
│   │   └── topics.go
│   └── visao_geral_do_curso
│       ├── overview.yml
│       └── topics.go
├── pkg
│   ├── format
│   │   ├── helpme.go
│   │   ├── menu_options.go
│   │   ├── outline_topic.go
│   │   ├── overview.go
│   │   ├── questionnaire.go
│   │   ├── resolucao_exercicios.go
│   │   └── section.go
│   └── reader
│       └── overview.go
└── tree.log

23 directories, 86 files
```

## TO-DO

- [ ] Refatorar o código para o uso de YAML para armazenar os tópicos.
- [ ] Implementar os tópicos restantes do curso.
- [ ] Implementar os exercícios restantes do curso.
- [ ] Implementar os testes unitários.
- [ ] Implementar o menu de overview por capítulo.
- [ ] Melhorar a documentação do projeto.
- [ ] Reorganizar a documentação em diretórios por capítulo.
- [ ] Melhorar a documentação das funções e métodos.
- [ ] Melhorar a apresentação da resolução dos exercícios.
- [ ] Implementar a funcionalidade de busca.
- [ ] Padronizar a linguagem usada para escrever o código, en-US.
- [ ] Implementar função de log para registrar as ações do usuário.

## Licença

Este projeto é licenciado sob a licença MIT. Consulte o arquivo [LICENSE](LICENSE) para obter detalhes.
