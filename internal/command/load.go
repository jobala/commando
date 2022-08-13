package command

import (
	"reflect"
	"strings"

	"github.com/akamensky/argparse"
)

/*
    Test add command
**/
func AddCommand[T any](cmdName string, handlerFunc func(args T), parser *argparse.Command) *argparse.Command {
	// create command,
	// Reflect on T to add arguments
	// Add command

	cmd := parser.NewCommand(cmdName, "A description")
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

func (c *UserCmd[T]) Load(cmdGrp *CommandGroup) *CommandGroup {
	AddCommand(c.Name, c.Handler, cmdGrp.Parent)
	return cmdGrp
}

func (cg *CommandGroup) WithCommand(cmd UserCmdr) *CommandGroup {
	cmd.Load(cg)
	return cg
}

type UserCmdr interface {
	Load(*CommandGroup) *CommandGroup
}

type CommandGroup struct {
	Parent *argparse.Command
}

type UserCmd[T any] struct {
	Name    string
	Handler func(args T)
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
