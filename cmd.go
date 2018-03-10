package main

import (
	"fmt"

	"github.com/atotto/clipboard"

	"github.com/urfave/cli"
)

// Emoji unicode
const (
	KeyEmoji = "\U0001f511"
	BoxEmoji = "\U0001f4e6"
)

func buildCommands() []cli.Command {
	commands := []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List all key-value pairs",
			Action: func(c *cli.Context) error {
				pairs, err := getAll()
				if err != nil {
					fmt.Println(err.Error())
					return nil
				}

				for k, v := range pairs {
					fmt.Printf("%s ) %s\n%s ) %s \n", KeyEmoji, k, BoxEmoji, v)
				}
				return nil
			},
		},
		{
			Name:  "get",
			Usage: "Get a value of a specific key",
			Action: func(c *cli.Context) error {
				if len(c.Args()) < 1 {
					fmt.Println("You must give a key")
					return nil
				}

				k := c.Args()[0]
				if v, err := get(k); err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println(v)
				}
				return nil
			},
		},
		{
			Name:  "set",
			Usage: "Set a key-value pair",
			Action: func(c *cli.Context) error {
				if len(c.Args()) < 2 {
					fmt.Println("You must give a key-value pair")
					return nil
				}

				k, v := c.Args()[0], c.Args()[1]
				if err := set(k, v); err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Printf("%s:%s is set\n", k, v)
				}
				return nil
			},
		},
		{
			Name:  "del",
			Usage: "Delete a key",
			Action: func(c *cli.Context) error {
				if len(c.Args()) < 1 {
					fmt.Println("You must give a key to delete")
					return nil
				}

				k := c.Args()[0]
				if err := del(k); err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Printf("%s is deleted\n", k)
				}
				return nil
			},
		},
		{
			Name:  "cp",
			Usage: "Copy a value of a specific key to clipboard",
			Action: func(c *cli.Context) error {
				if len(c.Args()) < 1 {
					fmt.Println("You must give a key to copy")
					return nil
				}

				k := c.Args()[0]
				if v, err := get(k); err != nil {
					fmt.Println(err.Error())
				} else {
					clipboard.WriteAll(string(v))
					fmt.Printf("%s is copied!\n", v)
				}
				return nil
			},
		},
	}
	return commands
}
