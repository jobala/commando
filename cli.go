package commando

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/jobala/commando/internal/command"
)

/**
TODO
    1. Increase the number of supported arguments
    2. Help Information
    3. Improve command table lookup time complexity
*/
func NewCli(name, description string) *cli {
	return &cli{
		parser: argparse.NewParser(name, description),
	}
}

func (c *cli) Invoke(args []string) {
	err := c.parser.Parse(os.Args)
	if err != nil {
		fmt.Println(err)
	}

	for _, cmdItem := range command.CommandTable {
		if cmdItem.Cmd.Happened() {
			cmdItem.Executor.Execute()
		}
	}
}

func (c *cli) NewCommandGroup(name, desc string) *command.CommandGroup {
	return &command.CommandGroup{Group: c.parser.NewCommand(name, desc)}
}

func Command[T any](cmdName string, handlerFunc func(args T)) command.UserCmdr {
	return &command.UserCmd[T]{
		Name:    cmdName,
		Handler: handlerFunc,
	}
}

type cli struct {
	parser *argparse.Parser
}
