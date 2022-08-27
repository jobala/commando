package command

import (
	"reflect"
	"strings"

	"github.com/akamensky/argparse"
)

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

func getCmdArgsFromHandler[T any](handlerFunc Handler[T]) map[string]string {
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
