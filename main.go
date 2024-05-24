// Aprenda Go: https://youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&si=_JIbmByhwYvHdJAr
package main

import (
	"fmt"
	"os"

	"github.com/fabianoflorentino/aprendago/outline"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("\nÉ necessário passar um argumento para o programa")
		fmt.Println("\nExemplo: go run main.go --bem-vindo")
		fmt.Print(outline.HELPME)
		return
	} else {
		outline.Menu(os.Args[1])
	}
}
