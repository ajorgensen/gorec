package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ajorgensen/gorec/gorec"
	"github.com/urfave/cli/v2"
)

func version() *cli.Command {
	return &cli.Command{
		Name: "version",
		Action: func(c *cli.Context) error {
			fmt.Println(gorec.Version)
			return nil
		},
	}
}

func main() {
	app := &cli.App{
		Name:  "gorec",
		Usage: "Make http requests with ease",
		Commands: []*cli.Command{
			version(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
