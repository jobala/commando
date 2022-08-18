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

func addCommandToGroup[T any](cmdName string, handlerFunc Handler[T], group *argparse.Command) {
	cmd := createNewCommand(cmdName, "A command", group)
	args := getCmdArgsFromHandler(handlerFunc)
	cmdArgs := addArgsToCmd(args, cmd)

	cliCmd := Command[T]{
		args:        cmdArgs,
		handlerFunc: handlerFunc,
	}

	addCmdToTable[T](cmd, cliCmd)
}

func createNewCommand(name, description string, group *argparse.Command) *argparse.Command {
	return group.NewCommand(name, description)
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
