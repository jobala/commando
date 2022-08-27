package command

import (
	"github.com/akamensky/argparse"
)

func addCommandToGroup[T any](name, description string, handler func(args T)) CliCommand[T] {
	return CliCommand[T]{
		Name:        name,
		Description: description,
		Handler:     handler,
	}
}

func (c CliCommand[T]) Load(group string, parser *argparse.Parser) {
	subparser := getSubParser(group, parser)

	cmd := subparser.NewCommand(c.Name, c.Description)
	args := getCmdArgsFromHandler(c.Handler)
	cmdArgs := addArgsToCmd(args, cmd)

	CommandList = append(CommandList, CommandItem{
		Cmd: cmd,
		Executor: Command[T]{
			args:        cmdArgs,
			handlerFunc: c.Handler,
		},
	})
}

type Handler[T any] func(args T)

type CliCommand[T any] struct {
	Name        string
	Description string
	Handler     func(args T)
}

type Loader interface {
	Load(string, *argparse.Parser)
}
