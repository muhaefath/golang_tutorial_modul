package model

import (
	"fmt"

	"github.com/urfave/cli"
)

//go:generate ./newList.sh Person
type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%v (%v)", p.Name, p.Age)
}

func Info(app *cli.App) {
	app.Name = "Hercules"
	app.Usage = "Generate MP Connector File"
	app.Author = "Sirclo"
	app.Version = "1.0.0"
}

func Commands(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name:    "peppers",
			Aliases: []string{"p"},
			Usage:   "Add peppers to your pizza",
			Action: func(c *cli.Context) {

				//go:generate ./newList.sh Person
				fmt.Println("success")
			},
		},
		{
			Name:    "pineapple",
			Aliases: []string{"pa"},
			Usage:   "Add pineapple to your pizza",
			Action: func(c *cli.Context) {
				//go:generate ./newList.sh Person
			},
		},
		{
			Name:    "cheese",
			Aliases: []string{"c"},
			Usage:   "Add cheese to your pizza",
			Action: func(c *cli.Context) {
				//go:generate ./newList.sh Person
			},
		},
	}
}