package main

import (
	"fmt"
	"log"
	"os"

	"github.com/muhaefath/golang_tutorial_modul/v2/modules/model"
	"github.com/urfave/cli"
)

var app = cli.NewApp()
var app2 = cli.NewApp()

//go:generate ./modules/model/newList.sh Person
func main() {
	model.Info(app)
	model.Commands(app)
	// Commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Commands() {
	app.Commands = []cli.Command{
		{
			Name:    "peppers",
			Aliases: []string{"a"},
			Usage:   "Add peppers to your pizza",
			Action: func(c *cli.Context) {
				//go:generate ./modules/model/newList.sh Person

				fmt.Println("success")
			},
		},
	}
}

type ClientFunction func() error

func TestFunction() error {

	fmt.Println("halo test function")
	return nil
}
