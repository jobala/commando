package command

import (
	"reflect"
	"strings"

	"github.com/akamensky/argparse"
)

func (cg *CommandGroup) WithCommand(cmd UserCmdr) *CommandGroup {
	cmd.Add(cg)
	return cg
}

func (c *UserCmd[T]) Add(cmdGroup *CommandGroup) *CommandGroup {
	addCommandToGroup(c.Name, c.Handler, cmdGroup.Group)
	return cmdGroup
}

func addCommandToGroup[T any](cmdName string, handlerFunc Handler[T], group *argparse.Command) *argparse.Command {
	cmd := createNewCommand(cmdName, "A command", group)
	args := getCmdArgsFromHandler(handlerFunc)
	addArgsToCmd(args, cmd)

	cliCmd := Command[T]{
		args:        args,
		handlerFunc: handlerFunc,
	}

	addCmdToTable[T](cmd, cliCmd)
	return cmd
}

func createNewCommand(name, description string, group *argparse.Command) *argparse.Command {
	return group.NewCommand(name, description)
}

func addArgsToCmd(args map[string]any, cmd *argparse.Command) {
	for field, dataType := range args {
		switch dataType {
		case "int":
			args[field] = cmd.Int("", field, nil)
		case "string":
			args[field] = cmd.String("", field, nil)
		}
	}
}

func getCmdArgsFromHandler[T any](handlerFunc Handler[T]) map[string]any {
	var argStruct T
	args := make(map[string]any)

	argReflection := reflect.TypeOf(argStruct)
	for i := 0; i < argReflection.NumField(); i++ {
		curField := argReflection.Field(i)
		field := strings.ToLower(curField.Name)
		args[field] = curField.Type.Name()
	}
	return args
}

func addCmdToTable[T any](cmd *argparse.Command, cliCmd CMD) {
	CommandTable = append(CommandTable, CommandItem{
		Cmd:      cmd,
		Executor: cliCmd,
	})
}

type CommandGroup struct {
	Group *argparse.Command
}

type UserCmd[T any] struct {
	Name    string
	Handler func(args T)
}

type UserCmdr interface {
	Add(*CommandGroup) *CommandGroup
}

type CommandItem struct {
	Cmd      *argparse.Command
	Executor CMD
}

type Handler[T any] func(args T)

var CommandTable = make([]CommandItem, 0)
