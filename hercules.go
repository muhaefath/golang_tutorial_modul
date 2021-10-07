package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()
var app2 = cli.NewApp()

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

type ClientFunction func(string) error
