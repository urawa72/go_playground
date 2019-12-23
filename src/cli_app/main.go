package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "translation-test"
	app.Usage = "translation on command"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {
		fmt.Println("Hello world in cli")
		return nil
	}

	app.Run(os.Args)
}
