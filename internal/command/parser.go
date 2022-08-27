package command

import (
	"reflect"
	"strings"

	"github.com/akamensky/argparse"
)

var subParsers = make(map[string]*argparse.Command)
var CommandList = []CommandItem{}

func getSubParser(group string, parser *argparse.Parser) *argparse.Command {
	if subParser, ok := subParsers[group]; ok {
		return subParser
	}

	groups := strings.Split(group, " ")

	grp := groups[0]

	var currCmd *argparse.Command
	if subParser, ok := subParsers[grp]; ok {
		currCmd = subParser
	} else {
		currCmd = parser.NewCommand(grp, "")

	}

	subParsers[grp] = currCmd

	for i := 1; i < len(groups); i++ {
		grp += " " + groups[i]
		currCmd = currCmd.NewCommand(grp, "")
		subParsers[grp] = currCmd
	}

	return subParsers[group]
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

type CommandItem struct {
	Cmd      *argparse.Command
	Executor CMD
}
