package command

import (
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
		subParsers[grp] = currCmd
	}

	for i := 1; i < len(groups); i++ {
		grp += " " + groups[i]
		currCmd = currCmd.NewCommand(groups[i], "")
		subParsers[grp] = currCmd
	}

	return subParsers[group]
}

type CommandItem struct {
	Cmd      *argparse.Command
	Executor CMD
}
