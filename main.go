// Aprenda Go: https://youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&si=_JIbmByhwYvHdJAr
package main

import (
	"os"

	outline "github.com/fabianoflorentino/aprendago/outline/menu"
)

func main() {
	if len(os.Args) < 2 {
		outline.PrintHelpMe()

		return
	} else {
		outline.Menu(os.Args[1:])
	}
}
