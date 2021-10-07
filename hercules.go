package main

import (
	"log"
	"os"

	"github.com/muhaefath/golang_tutorial_modul/v2/modules/model"
	"github.com/urfave/cli"
)

var app = cli.NewApp()
var app2 = cli.NewApp()

func main() {
	model.Info(app)
	model.Commands(app)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func NewOrder(webhook WebhookHandler) {

	webhook.NewOrder(nil, nil)
}

type ClientFunction func(string) error
