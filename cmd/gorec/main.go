package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ajorgensen/gorec/gorec"
	"github.com/joho/godotenv"
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

func send() *cli.Command {
	return &cli.Command{
		Name: "send",
		Action: func(c *cli.Context) error {
			var env map[string]string
			env, _ = godotenv.Read()

			filePath := c.Args().First()

			r, err := gorec.ParseFile(filePath, env)
			if err != nil {
				return err
			}

			resp, err := gorec.Do(r)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			fmt.Println(string(body))
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
			send(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
