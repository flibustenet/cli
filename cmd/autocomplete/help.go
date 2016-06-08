package autocomplete

import (
	"fmt"

	"github.com/Scalingo/codegangsta-cli"
)

func HelpAutoComplete(c *cli.Context) error {
	for cmd := range c.App.Commands {
		fmt.Println(c.App.Commands[cmd].Name)
	}

	return nil
}
