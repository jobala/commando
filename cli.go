package commando

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/jobala/commando/internal/command"
)

// NewCli creates a new Cli instance
func NewCli(name, description string) *cli {
	return &cli{
		parser:       argparse.NewParser(name, description),
		currentGroup: "",
	}
}

// NewCommandGroup creates a command group. Commands are added to a command group
//
//	NewCommandGroup("mammals cats")
//
// Will add $cli mammals cats  command group to your app
func (c *cli) NewCommandGroup(name string) *cli {
	c.currentGroup = name
	return c
}

// WithCommand adds a command to a command group
//
//	NewCommandGroup("mammals cats").WithCommand(commando.Command("lion", Handler))
//
// Will add $cli mammals cats lion --arg1 arg1 to your app
func (c *cli) WithCommand(cmd command.Loader) *cli {
	cmd.Load(c.currentGroup, c.parser)
	return c
}

// Command creates a command. The handlerFunc's args will be used as the command's arguments
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
	parser       *argparse.Parser
	currentGroup string
}
