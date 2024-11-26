/*
Main function to start the application and call the menu options based on the arguments
passed by the user.
If no arguments are passed, the HelpMe function is called to show the available options.
If arguments are passed, the Options function is called to execute the requested option.
  - os.Args[1:] is used to pass all arguments except the first one, which is the name of the program.
    This way, the Options function receives only the arguments that are relevant to the requested option.
    For example, if the user types "--outline", the Options function will receive only "--outline".
*/
package main

import (
	"os"

	"github.com/fabianoflorentino/aprendago/internal/menu"
)

func main() {
	if len(os.Args) < 2 {
		menu.HelpMe()

		return
	} else {
		menu.Options(os.Args[1:])
	}
}
