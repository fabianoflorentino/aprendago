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
    m.SetReady(80, 24)
    fmt.Println(m.View())
}
