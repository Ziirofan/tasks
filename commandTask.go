package main

import (
	"errors"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func commandTask() {
	app := &cli.App{
		Name:  "tasker",
		Usage: "A simple CLI program to manage your tasks",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					str := c.Args().First()
					if str == "" {
						return errors.New("Cannot add an empty task")
					}
					writeTask(1, str)
					return nil
				},
			},
			{
				Name:    "read",
				Aliases: []string{"r"},
				Usage:   "read a task by id",
				Action: func(c *cli.Context) error {
					str := c.Args().First()
					if str == "" {
						return errors.New("Cannot find without id")
					}
					displayTask(readTask(str))
					return nil
				},
			},
			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "update text of a task",
				Action: func(c *cli.Context) error {
					str := c.Args().First()
					txt := c.Args().Get(1)
					if str == "" {
						return errors.New("Cannot find without id")
					}
					updateTask(str, txt)
					return nil
				},
			},
			{
				Name:    "done",
				Aliases: []string{"d"},
				Usage:   "update text of a task",
				Action: func(c *cli.Context) error {
					str := c.Args().First()
					if str == "" {
						return errors.New("Cannot find without id")
					}
					doneTask(str)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
