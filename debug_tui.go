package main

import (
    "fmt"
    "github.com/fabianoflorentino/aprendago/internal/chapter"
    "github.com/fabianoflorentino/aprendago/internal/tui"
)

func main() {
    chapters := chapter.All()
    if len(chapters) == 0 {
        fmt.Println("no chapters")
        return
    }
    m := tui.NewModel(chapters)
    m.width = 80
    m.height = 24
    m.ready = true
    fmt.Print(m.View())
}
