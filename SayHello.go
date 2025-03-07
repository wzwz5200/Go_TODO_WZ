package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)


func GreetCommand() *cli.Command {
	return &cli.Command{
		Name:  "greet",
		Usage: "Say hello to someone",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Value:   "World",
				Usage:   "Person to greet",
			},
		},
		Action: greetAction, // 绑定 Action
	}
}

func greetAction(ctx context.Context, c *cli.Command) error {
	name := c.String("name")
	fmt.Printf("Hello, %s!\n", name)
	return nil
}
