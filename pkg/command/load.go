package command

import (
	"reflect"
	"strings"

	"github.com/akamensky/argparse"
)

/*
    Test add command
**/
func AddCommand[T any](cmdName string, handlerFunc func(args T)) *argparse.Command {
	// create command,
	// Reflect on T to add arguments
	// Add command

	cmd := Parser.NewCommand(cmdName, "A description")
	parsedArgs := make(map[string]interface{})

	var f T

	ft := reflect.TypeOf(f)
	for i := 0; i < ft.NumField(); i++ {
		curField := ft.Field(i)

		switch curField.Type.Name() {
		case "int":
			parsedArgs[strings.ToLower(curField.Name)] = cmd.Int("", strings.ToLower(curField.Name), nil)
		case "string":
			parsedArgs[strings.ToLower(curField.Name)] = cmd.String("", strings.ToLower(curField.Name), nil)
		}
	}

	cmdObj := Command[T]{
		args:        parsedArgs,
		handlerFunc: handlerFunc,
	}

	CommandTable = append(CommandTable, CommandItem{
		Cmd:      cmd,
		Executor: cmdObj,
	})

	return cmd
}

type CommandGroup struct {
	parent *argparse.Command
}

func AddCommandGroup(group, description string) *CommandGroup {
	parent := Parser.NewCommand(group, description)
	return &CommandGroup{
		parent: parent,
	}
}

func (g CommandGroup) WithCommand(cmdName string, handlerFunc func(args any)) {
	AddCommand(cmdName, handlerFunc)
}

type CMD interface {
	Execute()
}

type CommandItem struct {
	Cmd      *argparse.Command
	Executor CMD
}

var Parser = argparse.NewParser("mg", "A new CLI")
var CommandTable = make([]CommandItem, 0)
