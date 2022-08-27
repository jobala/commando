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

	subGroups := strings.Split(group, " ")
	grp := subGroups[0]

	var currCmd *argparse.Command

	if subParser, ok := subParsers[grp]; ok {
		currCmd = subParser
	} else {
		currCmd = parser.NewCommand(grp, "")
		subParsers[grp] = currCmd
	}

	for i := 1; i < len(subGroups); i++ {
		grp += " " + subGroups[i]
		currCmd = currCmd.NewCommand(subGroups[i], "")
		subParsers[grp] = currCmd
	}

	return subParsers[group]
}
