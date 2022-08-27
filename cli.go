package commando

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/jobala/commando/internal/command"
)

func NewCli(name, description string) *cli {
	return &cli{
		parser:       argparse.NewParser(name, description),
		currentGroup: "",
	}
}

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

func (c *cli) NewCommandGroup(name string) *cli {
	c.currentGroup = name
	return c
}

func (c *cli) WithCommand(cmd command.Loader) *cli {
	cmd.Load(c.currentGroup, c.parser)
	return c
}

func Command[T any](cmdName string, handlerFunc func(args T)) command.Loader {
	return command.CliCommand[T]{
		Name:    cmdName,
		Handler: handlerFunc,
	}
}

type cli struct {
	parser       *argparse.Parser
	currentGroup string
}

type cliCommand interface {
}
