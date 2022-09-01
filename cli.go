// Package commando provides utilities for creating command groups and adding Commands
// to a command group
package commando

import (
	"fmt"
	"os"

	"github.com/jobala/commando/internal/command"
)

// NewCli creates a new Cli instance
func NewCli(name, description string) *cli {
	return &cli{
		parser:       command.NewParser(name, description),
		currentGroup: "",
	}
}

// NewCommandGroup adds a command group to the CLI and returns an instance of CLI
// which is used by WithCommand to add commands to the command group.
//
//	NewCommandGroup("mammals cats")
//
// Will add $cli mammals cats  command group to your app
func (c *cli) NewCommandGroup(name string) *cli {
	c.currentGroup = name
	return c
}

// WithCommand adds a command to the command group to which it is chained and returns
// the CLI instance so that more commands can be added to the command group.
//
//	NewCommandGroup("mammals cats").WithCommand(commando.Command("lion", Handler))
//
// Will add $cli mammals cats lion --arg1 arg1 to your app
func (c *cli) WithCommand(cmd command.Loader) *cli {
	cmd.Load(c.currentGroup, c.parser)
	return c
}

// Command returns a loadable command which the CLI loads during execution
func Command[T any](cmdName string, handlerFunc func(args T)) command.Loader {
	return command.CliCommand[T]{
		Name:    cmdName,
		Handler: handlerFunc,
	}
}

// Invoke is the CLI's entry point
func (c *cli) Invoke(args []string) {
	err := c.parser.Parse(os.Args)
	if err != nil {
		fmt.Println(err)
	}

	for _, cmd := range command.CommandList {
		if cmd.Cmd.Happened() {
			cmd.Executor.Execute()
		}
	}
}

type cli struct {
	parser       *command.Parser
	currentGroup string
}
