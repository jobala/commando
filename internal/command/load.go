package command

import (
	"reflect"
	"strings"

	"github.com/akamensky/argparse"
)

// Load loads a commmand and flags into argparse
func (c CliCommand[T]) Load(group string, parser *argparse.Parser) {
	subparser := getSubParser(group, parser)

	cmd := subparser.NewCommand(c.Name, c.Description)
	args := getArgsFromHandler(c.Handler)
	cmdArgs := addArgsToCmd(args, cmd)

	cmdItem := CommandItem{
		Cmd: cmd,
		Executor: Command[T]{
			args:        cmdArgs,
			handlerFunc: c.Handler,
		},
	}

	addToCommandList(cmdItem)
}

func addToCommandList(cmdItem CommandItem) {
	CommandList = append(CommandList, cmdItem)
}

func getArgsFromHandler[T any](handlerFunc Handler[T]) map[string]string {
	var argStruct T
	args := make(map[string]string)

	argReflection := reflect.TypeOf(argStruct)
	for i := 0; i < argReflection.NumField(); i++ {
		curField := argReflection.Field(i)
		field := strings.ToLower(curField.Name)
		args[field] = curField.Type.Name()
	}

	return args
}

func addArgsToCmd(args map[string]string, cmd *argparse.Command) map[string]any {
	cmdArgs := make(map[string]any)

	for field, dataType := range args {
		switch dataType {
		case "int":
			cmdArgs[field] = cmd.Int("", field, nil)
		case "string":
			cmdArgs[field] = cmd.String("", field, nil)
		}
	}

	return cmdArgs
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

type CommandItem struct {
	Cmd      *argparse.Command
	Executor Executor
}
