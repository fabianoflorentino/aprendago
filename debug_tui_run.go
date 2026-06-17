package main

import (
    "fmt"
    "strings"
    "github.com/fabianoflorentino/aprendago/internal/chapter"
    "github.com/fabianoflorentino/aprendago/internal/tui"
    tea "github.com/charmbracelet/bubbletea"
)

func main() {
    chapters := chapter.All()
    if len(chapters) == 0 {
        fmt.Println("no chapters")
        return
    }
    m := tui.NewModel(chapters)
    m, _ = m.Update(tea.WindowSizeMsg{Width: 80, Height: 40})
    out := m.(tui.Model).View()
    lines := strings.Split(out, "\n")
    fmt.Println("line count:", len(lines))
    for i, line := range lines {
        fmt.Printf("%02d|%s\n", i+1, line)
    }
}
