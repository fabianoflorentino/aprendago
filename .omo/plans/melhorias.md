# Plano de Melhorias — aprendago CLI

## Priorização (P0 = mais impacto, menor esforço)

### P0 — Rápidas, alto retorno

| # | Tarefa | Esforço | Arquivos | Dependências |
|---|--------|---------|----------|--------------|
| 1 | Remover wrapper `ReadOverview`/`readOverview` | minutos | `pkg/reader/overview.go` | nenhuma |
| 2 | `aprendago cap <N> random` — tópico aleatório | ~30min | `cmd/cap.go` | nenhuma |
| 3 | Atualizar dependências (`go get -u ./...`) | ~15min | `go.mod`, `go.sum` | nenhuma |

### P1 — Médio retorno, esforço moderado

| # | Tarefa | Esforço | Arquivos | Dependências |
|---|--------|---------|----------|--------------|
| 4 | `aprendago search <term>` — busca textual | ~2h | `cmd/search.go` (novo), `cmd/register.go`, `internal/chapter/` | P0#3 |
| 5 | `golangci-lint` no CI | ~1h | `.golangci.yml`, `.github/workflows/lint.yml` | nenhuma |
| 6 | Colored output (`fatih/color`) | ~2h | `pkg/format/`, `cmd/` | P0#3 |
| 7 | GoReleaser + release automático | ~2h | `.goreleaser.yaml`, `.github/workflows/release.yml` | nenhuma |

### P2 — Maior esforço, médio retorno

| # | Tarefa | Esforço | Arquivos | Dependências |
|---|--------|---------|----------|--------------|
| 8 | Refatorar `fmt.Scan` → `io.Reader` injetável | ~3h | `pkg/format/questionnaire.go`, `internal/exercicios_ninja_nivel_1/`, `internal/exercicios_ninja_nivel_2/` | nenhuma |
| 9 | Refatorar `canais` para testável (context.Context) | ~4h | `internal/canais/` | nenhuma |

### P3 — Alto impacto, maior investimento

| # | Tarefa | Esforço | Arquivos | Dependências |
|---|--------|---------|----------|--------------|
| 10 | **TUI (bubbletea)** — navegação interativa do curso | ~8h | `internal/tui/` (novo pacote), `cmd/tui.go` (novo) | `github.com/charmbracelet/bubbletea`, `github.com/charmbracelet/bubbles`, `github.com/charmbracelet/lipgloss` |

---

## 1. Remover wrapper ReadOverview

**Estado atual:** `reader.ReadOverview()` exportado chama `readOverview()` não exportado com exatamente a mesma assinatura e corpo. Wrapper sem valor.

**Ação:** Matar a versão não exportada, renomear a exportada de `ReadOverview` para só ela existir.

**Verificação:** `go build ./...`, `go test ./...`, `go vet ./...`

---

## 2. `aprendago cap <N> random`

**Especificação:** 
```shell
aprendago cap 6 random
# → executa um tópico aleatório do capítulo 6
```

**Mudanças em `cmd/cap.go`:**
- Adicionar `case "random"` no switch de args[1]
- Usar `math/rand` para sortear um índice de `ch.Topics()`
- Chamar `ch.ExecTopic(topics[idx])` com o tópico sorteado

**Casos de borda:**
- Capítulo sem tópicos → erro claro
- Capítulo não encontrado já tratado pelo `chapter.Get()` existente

**Testes:**
- `TestCapRandomInvalidChapter` — capítulo inexistente → erro
- `TestCapRandom` — executar com output capture, verificar que imprime algo

---

## 3. `aprendago search <term>`

**Especificação:**
```shell
aprendago search goroutine
# → Capítulo X: Nome do Tópico
#   (primeiras linhas do texto)
```

**Implementação:**
- Novo arquivo `cmd/search.go` com cobra command `searchCmd`
- Registrar em `cmd/root.go`: `rootCmd.AddCommand(searchCmd)`
- Lógica: `chapter.All()` → `reader.ReadOverview(rootDir)` → `Description.Sections` → filtrar por `strings.Contains(text, term)` OU `strings.Contains(title, term)`
- Exibir resultados com capítulo + título da seção

**Testes:**
- `TestSearchFound` — busca termo que existe → 1+ resultados
- `TestSearchNotFound` — busca termo inexistente → "Nenhum resultado"
- `TestSearchEmptyTerm` — termo vazio → erro

---

## 4. Atualizar dependências

```shell
go get -u ./...
go mod tidy
```

Verificar breaking changes nos testes.

---

## 5. `golangci-lint` no CI

**Config `.golangci.yml`:**
```yaml
linters:
  enable:
    - errcheck
    - gosimple
    - govet
    - ineffassign
    - staticcheck
    - unused
    - revive
```

**Workflow `.github/workflows/lint.yml`:**
```yaml
name: Lint
on: [push, pull_request]
jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: golangci/golangci-lint-action@v6
```

**Meta:** Não quebrar CI. Se houver warnings, decidir se corrige ou silencia no config.

---

## 6. Colored output (`fatih/color`)

**Dependência:** `github.com/fatih/color`

**Mudanças:**
- `cmd/caps.go` — verde para títulos dos capítulos
- `cmd/cap.go` — azul para títulos de tópicos, amarelo para erros
- `cmd/outline.go` — cores por nível
- `pkg/format/overview.go` — template com cores

**Riscos:**
- Quebra capture de output em testes existentes (que comparam string exata)
- Solução: usar flag `--no-color` ou `NO_COLOR` env var nos testes

**Testes:** Ajustar output esperado nos testes existentes OU usar `NO_COLOR=1` nos testes.

---

## 7. GoReleaser + Release automático

**`.goreleaser.yaml`:**
```yaml
before:
  hooks:
    - go mod tidy
builds:
  - main: ./cmd/aprendago/main.go
    binary: aprendago
    goos: [linux, darwin, windows]
    goarch: [amd64, arm64]
```

**Workflow `.github/workflows/release.yml`:**
- Trigger: push de tag `v*`
- Usar `goreleaser-action`
- Publicar no GitHub Releases + DockerHub

---

## 8. Refatorar `fmt.Scan` → `io.Reader` injetável

**Problema:** `AnswerService.Collect()` usa `fmt.Scan` direto → impossível testar sem stdin real.

**Solução:** Adicionar campo `Reader io.Reader` em `AnswerService`, usar `bufio.NewScanner(s.Reader)` no lugar de `fmt.Scan`. Manter zero-value = `os.Stdin` para não quebrar produção.

```go
type AnswerService struct {
    Reader io.Reader
}

func (as AnswerService) reader() io.Reader {
    if as.Reader == nil {
        return os.Stdin
    }
    return as.Reader
}
```

**Arquivos afetados:**
- `pkg/format/questionnaire.go`
- `internal/exercicios_ninja_nivel_1/resolution_exercises.go`
- `internal/exercicios_ninja_nivel_2/resolution_exercises.go`

**Testes:** Injetar `strings.NewReader("a\nb\nc\n")`, capturar output, verificar `Answer` retornado.

---

## 9. Refatorar `canais` para ser testável

**Problema:** `UsingConvergeString`, `UsingDivergenceWithFunc`, etc. usam loops infinitos ou `time.Sleep` — testar exige timeout externo.

**Solução:** Adicionar parâmetro `ctx context.Context` e usar `select` com `<-ctx.Done()`. Garantir que as funções terminem quando o contexto for cancelado.

**Manter backward compatibility:** Versão sem `ctx` chama `context.Background()` internamente.

**Testes:** `ctx, cancel := context.WithTimeout(...)`, chamar função, cancelar, verificar que retornou sem leak de goroutine.

## 10. TUI (Terminal User Interface) com bubbletea

**Estado atual:** CLI puro baseado em cobra — toda interação exige digitar comandos e subcomandos. Navegação entre capítulos/tópicos é sequencial, sem visão exploratória.

**Proposta:** Adicionar `aprendago tui` que abre uma interface interativa no terminal com split view.

### Mockup conceitual

```
┌──────────────────────────────────────────────────┐
│  aprendago  (q: sair  ↑↓: navegar  /: buscar)    │
├─────────────────┬────────────────────────────────┤
│  Capítulos      │  Conteúdo                      │
│                 │                                │
│  ▶ 1. Visão Ge… │  Bem-vindo!                    │
│    2. Variáveis… │  - Bem vindo ao curso!         │
│    3. Exercício… │  - Go foi criada por gente…    │
│    4. Fundament… │  - É uma linguagem que vem…    │
│    5. Exercício… │  - Nesse curso você vai…       │
│    6. Fluxo de… │                                │
│    7. Exercício… │                                │
│    8. Agrupamen… │                                │
│    9. Exercício… │                                │
│   10. Structs   │                                │
│   11. Exercício… │                                │
│   12. Funções   │                                │
│   13. Exercício… │                                │
│   14. Ponteiros │                                │
│   15. Exercício… │                                │
│   16. Aplicações│                                │
│                 │                                │
│  [/] busca      │                                │
└─────────────────┴────────────────────────────────┘
```

### Arquitetura do pacote `internal/tui/`

| Arquivo | Responsabilidade |
|---------|-----------------|
| `model.go` | Model bubbletea: estado global (capítulo selecionado, tópicos carregados, conteúdo visível, dimensões da janela) |
| `update.go` | Gerenciamento de mensagens: teclas (↑↓←Enter/q,), redimensionamento, commands assíncronos (carregar overview.yml) |
| `view.go` | Renderização do layout dividido: lista à esquerda, viewport à direita, barra de busca no rodapé |
| `styles.go` | Definições lipgloss: cores, bordas arredondadas, dimensões dos painéis, padding |
| `chapters.go` | Componente `bubbles/list` para listar capítulos, com indicador de seleção e scroll |
| `topics.go` | Sub-lista de tópicos do capítulo selecionado (expande/colapsa com Enter) |
| `content.go` | `bubbles/viewport` com scroll do conteúdo formatado via `section.New().Format()` |

### Comando cobra

**`cmd/tui.go`:**
```go
var tuiCmd = &cobra.Command{
    Use:   "tui",
    Short: "Modo interativo com navegação visual do curso",
    RunE: func(cmd *cobra.Command, args []string) error {
        return runTUI()
    },
}
```

Registrar em `root.go`: `rootCmd.AddCommand(tuiCmd)`

### Fluxo de navegação

```
aprendago tui
  → bubbletea.NewProgram(model.InitialModel())
    → painel esquerdo: chapter.All() → bubbles/list
    → Enter no capítulo → Topics() → mostra tópicos no mesmo painel (indentados)
    → Enter no tópico → section.New().Format() via tea.Cmd assíncrono
      → resultado renderizado no viewport (painel direito)
    → ↑↓ → navega entre itens
    → ← / Esc → volta da lista de tópicos para lista de capítulos
    → / → ativa barra de busca (filtra capítulos/tópicos por texto)
    → q / Ctrl+C → sair
    → redimensionamento → tea.WindowSizeMsg → ajusta painéis
```

### Reaproveitamento do código existente

Tudo que já existe é 100% reaproveitado — a TUI só adiciona navegação:

| Código existente | Uso na TUI |
|-----------------|-----------|
| `chapter.All()` | Preencher lista de capítulos |
| `chapter.Get(num)` | Obter capítulo ao selecionar |
| `ch.Topics()` | Listar tópicos ao expandir capítulo |
| `section.New(dir).Format(title)` | Carregar conteúdo no viewport |
| `reader.ReadOverview(dir)` | (usado internamente por Topics/Format) |
| `format.FormatOverview()` | (usado internamente por Format) |

### Testes

- **Teste de modelo:** Criar model, enviar mensagens de teclas, verificar transições de estado
- **Teste de carregamento:** Simular seleção de capítulo, verificar que `Topics()` foi chamado
- **Teste de busca:** Inicializar model com termo de busca, verificar filtro na lista
- **Snapshot testing:** Renderizar view em estado conhecido e comparar com snapshot (opcional)

### Dependências novas

```
github.com/charmbracelet/bubbletea v1.x
github.com/charmbracelet/bubbles    v0.x
github.com/charmbracelet/lipgloss   v1.x
```

### Observações

- **bubbletea** usa o modelo Elm (Model → Msg → Update → View) — curva de aprendizado pequena para quem conhece React/Elm
- **bubbles** fornece componentes prontos: `list`, `viewport`, `textinput`, `spinner` (para loading enquanto carrega overview.yml)
- **lipgloss** lida com estilos, cores 16/256/truecolor, adapta ao terminal
- A TUI **não substitui** o CLI original — é um comando adicional. Scripts e CI continuam usando `caps`/`cap`/`outline`

---

## Risco e Mitigações

| Risco | Probabilidade | Mitigação |
|-------|-------------|-----------|
| GoReleaser quebrar release manual atual | Baixa | Manter workflow atual como fallback |
| `fatih/color` quebrar testes de output | Alta | Usar `NO_COLOR` ou `—no-color` nos testes |
| `go get -u` atualizar versão major quebrando API | Média | `go get -u -m` pacote a pacote, `go mod verify` |
| `canais` com goroutine leak em teste | Média | Usar `t.Cleanup(cancel)` + `leakcheck` |
| bubbletea v1 com breaking changes | Média | Fixar versão no `go.mod`, testar upgrade separadamente |
| TUI não renderizar bem em terminais muito pequenos (<80 colunas) | Média | Viewport mínimo com fallback para layout vertical em tela estreita |

## Ordem sugerida de implementação

```
P0 (rápidas) → P1 (funcionalidades CLI) → P3 (TUI) → P2 (refatorações internas)
```

**Justificativa:** P0 desbloqueia a base. P1 melhora a experiência do CLI que existe hoje. A TUI (P3) é a feature de maior impacto e deve vir antes das refatorações internas (P2) porque ela reaproveita o código atual — se refatorarmos primeiro e depois fizermos a TUI, corremos risco de refatorar código que poderia nem precisar mudar. O contrário (TUI primeiro) revela exatamente quais interfaces precisam de refatoração.
