# Aprenda Go

[![Go Report Card](https://goreportcard.com/badge/github.com/fabianoflorentino/aprendago)](https://goreportcard.com/report/github.com/fabianoflorentino/aprendago) [![Build, Publish, Tag and Release](https://github.com/fabianoflorentino/aprendago/actions/workflows/ci.yml/badge.svg)](https://github.com/fabianoflorentino/aprendago/actions/workflows/ci.yml) [![CodeQL](https://github.com/fabianoflorentino/aprendago/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/fabianoflorentino/aprendago/actions/workflows/github-code-scanning/codeql) [![Trivy vulnerability scan](https://github.com/fabianoflorentino/aprendago/actions/workflows/trivy.yml/badge.svg)](https://github.com/fabianoflorentino/aprendago/actions/workflows/trivy.yml) [![Dependabot Updates](https://github.com/fabianoflorentino/aprendago/actions/workflows/dependabot/dependabot-updates/badge.svg)](https://github.com/fabianoflorentino/aprendago/actions/workflows/dependabot/dependabot-updates)

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

## Usando o CLI

O CLI oferece três subcomandos principais:

### `aprendago caps` — Lista capítulos disponíveis

```shell
go run cmd/aprendago/main.go caps
```

```
Capítulos do Curso
  aprendago cap 1 topics    Visão Geral do Curso
  aprendago cap 2 topics    Variáveis, Valores e Tipos
  aprendago cap 3 topics    Exercícios Ninja: Nível 1
  ...
```

### `aprendago cap <N>` — Acessa um capítulo

```shell
go run cmd/aprendago/main.go cap 1 topics
```

```
Visão Geral do Curso
  Bem-vindo!
  Por que Go?
  Sucesso
  Recursos
  Como esse curso funciona
```

Subcomandos do `cap`:
| Uso | Descrição |
|---|---|
| `aprendago cap <N>` | Mostra o overview do capítulo |
| `aprendago cap <N> topics` | Lista os tópicos do capítulo |
| `aprendago cap <N> overview` | Mostra o conteúdo completo do capítulo |
| `aprendago cap <N> "<tópico>"` | Mostra um tópico específico |

### `aprendago outline` — Outline completo do curso

```shell
go run cmd/aprendago/main.go outline
```

```bash
Visão Geral do Curso
  Bem-vindo!
  Por que Go?
  ...
```

### `aprendago --help` — Ajuda

```shell
go run cmd/aprendago/main.go --help
```

```bash
CLI para o curso Aprenda Go

Usage:
  aprendago [flags]
  aprendago [command]

Available Commands:
  cap         Acessa um capítulo do curso
  caps        Lista capítulos disponíveis
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  outline     Exibe o outline completo do curso

Flags:
  -h, --help   help for aprendago

Use "aprendago [command] --help" for more information about a command.
```

## Fluxo de Funcionamento

```mermaid
flowchart TD
    CLI["aprendago &lt;comando&gt;"] --> Cobra["cobra rootCmd.Execute()"]
    Cobra --> Caps["caps"]
    Cobra --> Cap["cap &lt;N&gt;"]
    Cobra --> Outline["outline"]
    Cobra --> Help["--help"]

    Caps --> CapsList["chapter.All()"]
    CapsList --> CapsPrint["Exibe lista de capítulos"]

    Cap --> CapParse["parseia args[0] → int"]
    CapParse --> CapGet["chapter.Get(num)"]
    CapGet --> CapArgs{"len(args) > 1?"}
    CapArgs -->|"topics"| CapTopics["printChapterTopics(ch)"]
    CapTopics --> TopicsRead["ch.Topics()"]
    TopicsRead --> TopicsRender["Lista títulos das seções"]
    CapArgs -->|"overview"| CapOverview["ch.Overview()"]
    CapOverview --> OverviewLoop["Para cada tópico → ch.ExecTopic()"]
    CapArgs -->|"&lt;tópico&gt;"| CapTopic["ch.ExecTopic(args[1])"]
    CapTopic --> SectionFmt["section.New(rootDir).Format(title)"]
    SectionFmt --> SectionRead["reader.ReadSection()"]
    SectionFmt --> OverviewFmt["format.FormatOverview()"]
    OverviewLoop --> SectionFmt

    Outline --> OutlineAll["chapter.All()"]
    OutlineAll --> OutlineEach["Para cada capítulo"]
    OutlineEach --> TopicsRead

    style CapsList fill:#3b82f6,color:#fff
    style CapGet fill:#3b82f6,color:#fff
    style SectionRead fill:#22c55e,color:#fff
    style OverviewFmt fill:#22c55e,color:#fff
```

## Arquitetura do Código

```mermaid
flowchart TD
    subgraph cmd["cmd/ — Cobra Commands"]
        Main["aprendago/main.go<br/>Execute()"]
        Root["root.go<br/>rootCmd + init()"]
        Cap["cap.go<br/>cap &lt;N&gt;"]
        Caps["caps.go<br/>caps"]
        Outline["outline.go<br/>outline"]
        Register["register.go<br/>blank imports → init()"]
    end

    subgraph internal["internal/ — Core Logic"]
        Chapter["chapter/<br/>Chapter type + registry"]
        ChapterTypes["types.go<br/>Overview(), ExecTopic(), Topics()"]
        Registry["registry.go<br/>Register(), All(), Get()"]
        ChapterPkgs["<capítulo>/*/<br/>27 packages com:<br/>chapter.go (init) + overview.yml"]
    end

    subgraph pkg["pkg/ — Shared Utilities"]
        Section["section/<br/>Section.New().Format()"]
        Reader["reader/<br/>YAML parsing"]
        Format["format/<br/>FormatOverview() template"]
        Output["output/<br/>Test helpers"]
        Trim["trim/<br/>String utils"]
    end

    Main --> Root
    Root --> Cap
    Root --> Caps
    Root --> Outline

    Cap --> Chapter
    Caps --> Chapter
    Outline --> Chapter

    ChapterTypes --> Section
    Chapter --> Registry
    Registry -.->|"init()"| ChapterPkgs
    Register -.-> ChapterPkgs

    Section --> Reader
    Section --> Format
    Reader --> Format
```

## Estrutura do Projeto

```shell
.
├── cmd/                    # CLI entrypoint and cobra commands
│   ├── aprendago/main.go
│   ├── cap.go              # aprendago cap <N> [topics|overview|<topic>]
│   ├── caps.go             # aprendago caps
│   ├── outline.go          # aprendago outline
│   ├── register.go         # blank imports for chapter self-registration
│   └── root.go             # cobra root command
├── internal/
│   ├── chapter/            # Chapter type, registry, and helpers
│   └── <capitulo>/         # 27 chapter packages (overview.yml + chapter.go)
├── pkg/                    # Shared utilities
│   ├── format/             # Output formatting (overview, questionnaire)
│   ├── output/             # Output helpers
│   ├── reader/             # YAML overview reader
│   ├── section/            # Section formatting from overview.yml
│   └── trim/               # String trimming
├── Makefile
├── Dockerfile
├── docker-compose.yml
└── README.md
```

## Licença

Este projeto é licenciado sob a licença MIT. Consulte o arquivo [LICENSE](LICENSE) para obter detalhes.
