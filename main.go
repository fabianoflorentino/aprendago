// Aprenda Go: https://youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&si=_JIbmByhwYvHdJAr
package main

import (
	"fmt"
	"os"

	outline "github.com/fabianoflorentino/aprendago/outline/menu"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("\nÉ necessário passar um argumento para o programa")
		fmt.Println("\nExemplo: go run main.go --bem-vindo")
		outline.PrintHelpMe()

		return
	} else {
		outline.Menu(os.Args[1])
	}
}
