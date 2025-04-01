# Aprenda Go

[![Build, Publish, Tag and Release](https://github.com/fabianoflorentino/aprendago/actions/workflows/ci.yml/badge.svg)](https://github.com/fabianoflorentino/aprendago/actions/workflows/ci.yml) [![CodeQL](https://github.com/fabianoflorentino/aprendago/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/fabianoflorentino/aprendago/actions/workflows/github-code-scanning/codeql) [![Trivy vulnerability scan](https://github.com/fabianoflorentino/aprendago/actions/workflows/trivy.yml/badge.svg)](https://github.com/fabianoflorentino/aprendago/actions/workflows/trivy.yml) [![Dependabot Updates](https://github.com/fabianoflorentino/aprendago/actions/workflows/dependabot/dependabot-updates/badge.svg)](https://github.com/fabianoflorentino/aprendago/actions/workflows/dependabot/dependabot-updates)

[![Linux](https://github.com/fabianoflorentino/aprendago/actions/workflows/linux.yml/badge.svg)](https://github.com/fabianoflorentino/aprendago/actions/workflows/linux.yml) [![macOS](https://github.com/fabianoflorentino/aprendago/actions/workflows/macos.yml/badge.svg)](https://github.com/fabianoflorentino/aprendago/actions/workflows/macos.yml) [![Windows](https://github.com/fabianoflorentino/aprendago/actions/workflows/windows.yml/badge.svg)](https://github.com/fabianoflorentino/aprendago/actions/workflows/windows.yml)

[![Go Reference](https://pkg.go.dev/badge/github.com/fabianoflorentino/aprendago.svg)](https://pkg.go.dev/github.com/fabianoflorentino/aprendago)
[![DockerHub](https://img.shields.io/badge/Docker-hub-blue)](https://hub.docker.com/r/fabianoflorentino/aprendago/tags)

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

  --cap=1 --overview    Visão Geral do Curso
  --cap=2 --overview    Variáveis, Valores &amp; Tipos
  --cap=3 --overview    Exercícios Ninja Nível 1
  --cap=4 --overview    Fundamentos da Programação
  --cap=5 --overview    Exercícios Ninja Nível 2
  --cap=6 --overview    Fluxo de Controle
  --cap=7 --overview    Exercícios Ninja Nível 3
  --cap=8 --overview    Agrupamento de Dados
  --cap=9 --overview    Exercícios Ninja Nível 4
  --cap=10 --overview   Structs
  --cap=11 --overview   Exercícios Ninja Nível 5
  --cap=12 --overview   Funções
  --cap=13 --overview   Exercícios Ninja Nível 6
  --cap=14 --overview   Ponteiros
  --cap=15 --overview   Exercícios Ninja Nível 7
  --cap=16 --overview   Aplicações
  --cap=17 --overview   Exercícios Ninja Nível 8
  --cap=18 --overview   Concorrência
  --cap=19 --overview   Seu Ambiente de Desenvolvimento
  --cap=20 --overview   Exercícios Ninja Nível 9
  --cap=21 --overview   Canais
  --cap=22 --overview   Exercícios Ninja Nível 10
  --cap=23 --overview   Tratamento de Erro
  --cap=24 --overview   Exercícios Ninja Nível 11
  --cap=25 --overview   Documentação
  --cap=26 --overview   Exercícios Ninja Nível 12
  --cap=27 --overview   Teste e Benchmarking
```

### Ajuda

```shell
docker compose exec -it aprendago /bin/sh -c 'go run cmd/aprendago/main.go --help'
```

```shell
Learn GO


Uso: go run cmd/aprendago/main.go [opção]

Exemplo:
  go run cmd/aprendago/main.go --bem-vindo

Ajuda:

--outline  Exibe o outline completo do curso.
--help     Exibe a lista de todas as opções disponíveis.
--caps     Exibe a lista de capítulos disponíveis.

Capítulos do Curso

  --cap=1 --topics    Visão Geral do Curso
  --cap=2 --topics    Variáveis, Valores e Tipos
  --cap=3 --topics    Exercícios Ninja: Nível 1
  --cap=4 --topics    Fundamentos da Programação
  --cap=5 --topics    Exercícios Ninja: Nível 2
  --cap=6 --topics    Fluxo de Controle
  --cap=7 --topics    Exercícios Ninja: Nível 3
  --cap=8 --topics    Agrupamento de Dados
  --cap=9 --topics    Exercícios Ninja: Nível 4
  --cap=10 --topics   Structs
  --cap=11 --topics   Exercícios Ninja: Nível 5
  --cap=12 --topics   Funções
  --cap=13 --topics   Exercícios Ninja: Nível 6
  --cap=14 --topics   Ponteiros
  --cap=15 --topics   Exercícios Ninja: Nível 7
  --cap=16 --topics   Aplicações
  --cap=17 --topics   Exercícios Ninja: Nível 8
  --cap=18 --topics   Concorrência
  --cap=19 --topics   Seu Ambiente de Desenvolvimento
  --cap=20 --topics   Exercícios Ninja: Nível 9
  --cap=21 --topics   Canais
  --cap=22 --topics   Exercícios Ninja: Nível 10
  --cap=23 --topics   Tratamento de Erro
  --cap=24 --topics   Exercícios Ninja: Nível 11
  --cap=25 --topics   Documentação
  --cap=26 --topics   Exercícios Ninja: Nível 12
  --cap=27 --topics   Teste e Benchmarking

Outline do Curso por Capítulo

  --cap=1 --overview    Visão Geral do Curso
  --cap=2 --overview    Variáveis, Valores &amp; Tipos
  --cap=3 --overview    Exercícios Ninja Nível 1
  --cap=4 --overview    Fundamentos da Programação
  --cap=5 --overview    Exercícios Ninja Nível 2
  --cap=6 --overview    Fluxo de Controle
  --cap=7 --overview    Exercícios Ninja Nível 3
  --cap=8 --overview    Agrupamento de Dados
  --cap=9 --overview    Exercícios Ninja Nível 4
  --cap=10 --overview   Structs
  --cap=11 --overview   Exercícios Ninja Nível 5
  --cap=12 --overview   Funções
  --cap=13 --overview   Exercícios Ninja Nível 6
  --cap=14 --overview   Ponteiros
  --cap=15 --overview   Exercícios Ninja Nível 7
  --cap=16 --overview   Aplicações
  --cap=17 --overview   Exercícios Ninja Nível 8
  --cap=18 --overview   Concorrência
  --cap=19 --overview   Seu Ambiente de Desenvolvimento
  --cap=20 --overview   Exercícios Ninja Nível 9
  --cap=21 --overview   Canais
  --cap=22 --overview   Exercícios Ninja Nível 10
  --cap=23 --overview   Tratamento de Erro
  --cap=24 --overview   Exercícios Ninja Nível 11
  --cap=25 --overview   Documentação
  --cap=26 --overview   Exercícios Ninja Nível 12
  --cap=27 --overview   Teste e Benchmarking
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
│   │   ├── constants.go
│   │   ├── help.go
│   │   ├── menu.go
│   │   ├── menu_test.go
│   │   ├── overview.yml
│   │   ├── topic_test.go
│   │   └── topics.go
│   ├── aplicacoes
│   │   ├── constants.go
│   │   ├── examples.go
│   │   ├── help.go
│   │   ├── menu.go
│   │   ├── menu_test.go
│   │   ├── overview.yml
│   │   ├── topics.go
│   │   └── topics_test.go
│   ├── canais
│   │   ├── examples.go
│   │   ├── overview.yml
│   │   └── topics.go
│   ├── chapter
│   │   └── chapter.go
│   ├── concorrencia
│   │   ├── overview.yml
│   │   └── topics.go
│   ├── documentacao
│   │   ├── overview.yml
│   │   └── topics.go
│   ├── exercicios_ninja_nivel_1
│   │   ├── overview.yml
│   │   ├── resolution_exercises.go
│   │   ├── resolution_exercises_test.go
│   │   └── topics.go
│   ├── exercicios_ninja_nivel_10
│   │   ├── overview.yml
│   │   ├── resolution_exercises.go
│   │   ├── resolution_exercises_test.go
│   │   └── topics.go
│   ├── exercicios_ninja_nivel_11
│   │   ├── overview.yml
│   │   ├── resolution_exercises.go
│   │   ├── resolution_exercises_test.go
│   │   └── topics.go
│   ├── exercicios_ninja_nivel_12
│   │   ├── overview.yml
│   │   ├── resolution_exercises.go
│   │   ├── resolution_exercises_test.go
│   │   └── topics.go
│   ├── exercicios_ninja_nivel_2
│   │   ├── overview.yml
│   │   ├── resolution_exercises.go
│   │   ├── resolution_exercises_test.go
│   │   └── topics.go
│   ├── exercicios_ninja_nivel_3
│   │   ├── overview.yml
│   │   ├── resolution_exercises.go
│   │   ├── resolution_exercises_test.go
│   │   └── topics.go
│   ├── exercicios_ninja_nivel_4
│   │   ├── overview.yml
│   │   ├── resolution_exercises.go
│   │   ├── resolution_exercises_test.go
│   │   └── topics.go
│   ├── exercicios_ninja_nivel_5
│   │   ├── overview.yml
│   │   ├── resolution_exercises.go
│   │   ├── resolution_exercises_test.go
│   │   └── topics.go
│   ├── exercicios_ninja_nivel_6
│   │   ├── overview.yml
│   │   ├── resolution_exercises.go
│   │   ├── resolution_exercises_test.go
│   │   └── topics.go
│   ├── exercicios_ninja_nivel_7
│   │   ├── overview.yml
│   │   ├── resolution_exercises.go
│   │   ├── resolution_exercises_test.go
│   │   └── topics.go
│   ├── exercicios_ninja_nivel_8
│   │   ├── overview.yml
│   │   ├── resolution_exercises.go
│   │   ├── resolution_exercises_test.go
│   │   └── topics.go
│   ├── exercicios_ninja_nivel_9
│   │   ├── overview.yml
│   │   ├── resolution_exercise.go
│   │   ├── resolution_exercises_test.go
│   │   └── topics.go
│   ├── fluxo_de_controle
│   │   ├── overview.yml
│   │   └── topics.go
│   ├── funcoes
│   │   ├── overview.yml
│   │   ├── resolution_challange.go
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
│   ├── ponteiros
│   │   ├── overview.yml
│   │   └── topics.go
│   ├── seu_ambiente_de_desenvolvimento
│   │   ├── overview.yml
│   │   └── topics.go
│   ├── structs
│   │   ├── overview.yml
│   │   └── topics.go
│   ├── teste_benchmarking
│   │   ├── overview.yml
│   │   └── topics.go
│   ├── tratamento_de_erro
│   │   ├── overview.yml
│   │   └── topics.go
│   ├── variaveis_valores_tipos
│   │   ├── overview.yml
│   │   └── topics.go
│   └── visao_geral_do_curso
│       ├── overview.yml
│       └── topics.go
├── log
│   └── development.log
├── pkg
│   ├── base_content
│   │   ├── base_content.go
│   │   └── base_content_test.go
│   ├── format
│   │   ├── helpme.go
│   │   ├── menu_options.go
│   │   ├── overview.go
│   │   ├── questionnaire.go
│   │   ├── questionnaire_test.go
│   │   ├── resolucao_exercicios.go
│   │   └── section.go
│   ├── logger
│   │   ├── log.go
│   │   ├── logger.go
│   │   └── logger_test.go
│   ├── output
│   │   └── output.go
│   ├── reader
│   │   └── overview.go
│   ├── section
│   │   └── section.go
│   ├── topic
│   │   └── topic.go
│   └── trim
│       └── trim.go
└── tree.log

45 directories, 125 files
```

## Licença

Este projeto é licenciado sob a licença MIT. Consulte o arquivo [LICENSE](LICENSE) para obter detalhes.
