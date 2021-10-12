package modules

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

//go:generate ./newList.sh

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
	app.Version = "1.0.0"
}

func Commands(app *cli.App) {

	app.Commands = []*cli.Command{
		{
			Name:    "peppers",
			Aliases: []string{"p"},
			Usage:   "Add peppers to your pizza",
			Action: func(c *cli.Context) error {

				out, err := exec.Command("/bin/sh", "./modules/newList.sh").Output()
				if err != nil {
					fmt.Printf("error %s :", err)
				}
				fmt.Println(out)
				fmt.Println(" asas2")

				prg := "echo"

				arg1 := "there"
				arg2 := "are three"
				arg3 := "falcons"

				cmd := exec.Command(prg, arg1, arg2, arg3)
				stdout, err := cmd.Output()
				fmt.Println(string(stdout))
				return nil
			},
		},
		{
			Name:    "pineapple",
			Aliases: []string{"pa"},
			Usage:   "Add pineapple to your pizza",
			Action: func(c *cli.Context) error {
				//go:generate ./newList.sh Person
				return nil
			},
		},
		{
			Name:    "cheese",
			Aliases: []string{"c"},
			Usage:   "Add cheese to your pizza",
			Action: func(c *cli.Context) error {
				//go:generate ./newList.sh Person
				return nil
			},
		},
	}
}
