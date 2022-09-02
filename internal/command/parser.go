package command

import (
	"strings"

	"github.com/akamensky/argparse"
)

var CommandList = []CommandItem{}

func NewParser(name, description string) *Parser {
	return &Parser{
		name:        name,
		description: description,
		argParser:   argparse.NewParser(name, description),
		subParsers:  make(map[string]*argparse.Command),
	}
}

func (p *Parser) Parse(args []string) error {
	return p.argParser.Parse(args)
}

func (p *Parser) Usage(err error) string {
	return p.argParser.Usage(err)
}

func (p *Parser) getSubParser(group string) *argparse.Command {
	subParser, subParserExists := p.subParsers[group]

	if subParserExists {
		return subParser
	} else {
		p.createSubParserFor(group)
	}

	return p.subParsers[group]
}

func (p *Parser) createSubParserFor(group string) {
	subGroups := strings.Split(group, " ")

	var currCmd *argparse.Command
	currGroup := subGroups[0]

	if subParser, ok := p.subParsers[currGroup]; ok {
		currCmd = subParser
	} else {
		currCmd = p.argParser.NewCommand(currGroup, "")
		p.subParsers[currGroup] = currCmd
	}

	for i := 1; i < len(subGroups); i++ {
		currGroup += " " + subGroups[i]
		currCmd = currCmd.NewCommand(subGroups[i], "")
		p.subParsers[currGroup] = currCmd
	}
}

type Parser struct {
	argParser   *argparse.Parser
	subParsers  map[string]*argparse.Command
	name        string
	description string
}
