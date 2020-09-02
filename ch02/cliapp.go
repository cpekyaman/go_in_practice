package main

import (
	"fmt"
	"os"

	"log"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "CliApp"
	app.Usage = "A Cli application"

	cmdFlags := []cli.Flag{
		&cli.StringFlag{
			Name:    "name",
			Aliases: []string{"n"},
			Value:   "John Doe",
			Usage:   "Whom I am talking to",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "hello",
			Aliases: []string{"hl"},
			Usage:   "Say Hello",
			Action: func(c *cli.Context) error {
				name := c.String("name")
				fmt.Printf("Hello %s\n", name)
				return nil
			},
			Flags: cmdFlags,
		},
		{
			Name:    "bye",
			Aliases: []string{"by"},
			Usage:   "Say Goodbye",
			Action: func(c *cli.Context) error {
				name := c.String("name")
				fmt.Printf("Goodbye %s\n", name)
				return nil
			},
			Flags: cmdFlags,
		},
	}

	app.Action = func(ctx *cli.Context) error {
		name := ctx.String("name")
		fmt.Printf("Hello %s\n", name)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("Damn !")
	}
}
