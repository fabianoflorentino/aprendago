// Aprenda Go: https://youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&si=_JIbmByhwYvHdJAr
package main

import (
	"os"

	"github.com/fabianoflorentino/aprendago/outline/menu"
)

func main() {
	if len(os.Args) < 2 {
		menu.ShowHelpMe()

		return
	} else {
		menu.Options(os.Args[1:])
	}
}
