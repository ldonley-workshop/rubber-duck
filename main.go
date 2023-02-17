package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Virtual Rubber Duck",
		Usage: "A CLI tool to help you solve your problems by talking to a virtual rubber duck",
		Commands: []*cli.Command{
			{
				Name:    "ask",
				Usage:   "Ask the virtual rubber duck a question",
				Aliases: []string{"a"},
				Action:  askAction,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func askAction(c *cli.Context) error {
	// Get the question from the user
	question := c.Args().First()

	// Check if a question was provided
	if question == "" {
		fmt.Println("Please provide a question to ask the virtual rubber duck")
	} else {
		// Print the response from the virtual rubber duck to the standard output
		fmt.Fprintf(os.Stdout, "You asked: %s\n", question)
		fmt.Fprintln(os.Stdout, "The virtual rubber duck says: Quack quack! I think you should try doing it this way...")
	}

	return nil
}
