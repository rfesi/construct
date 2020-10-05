package main

import (
	"fmt"
	"os"

	"github.com/networkteam/construct"
)

func main() {
	app := construct.NewCliApp()
	err := app.Run(os.Args)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v", err)
	}
}
