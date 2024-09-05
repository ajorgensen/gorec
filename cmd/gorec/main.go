package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gorec",
		Usage: "Make http requests with ease",
		Action: func(*cli.Context) error {
			fmt.Println("Hello, gorec!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
