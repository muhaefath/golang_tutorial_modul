package main

import (
	"fmt"
	"log"
	"os"

	"github.com/muhaefath/golang_tutorial_modul/v2/modules"
	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()
var app2 = cli.NewApp()

//go:generate ./modules/newList.sh Person
func main() {
	modules.Info(app)
	modules.Commands(app)
	// Commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// e := InitializeString()
	// fmt.Println(e)
	// e.Start()
}

// func Commands() {
// 	app.Commands = []cli.Command{
// 		{
// 			Name:    "peppers",
// 			Aliases: []string{"a"},
// 			Usage:   "Add peppers to your pizza",
// 			Action: func(c *cli.Context) {
// 				//go:generate ./modules/model/newList.sh Person

// 				fmt.Println("success")
// 			},
// 		},
// 	}
// }

type ClientFunction func() error

func TestFunction() error {

	fmt.Println("halo test function")
	return nil
}
