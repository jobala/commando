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
	if group == "" {
		return &p.argParser.Command
	}

	subParser, subParserExists := p.subParsers[group]

	if subParserExists {
		return subParser
	} else {
		subGroups := strings.Split(group, " ")

		_, subParserExists := p.subParsers[subGroups[0]]
		if !subParserExists {
			initialSubparser := p.argParser.NewCommand(subGroups[0], "")
			p.subParsers[subGroups[0]] = initialSubparser
		}

		p.createSubParserFor(subGroups, nil, "", 0)
	}
	return p.subParsers[group]
}

func (p *Parser) createSubParserFor(groups []string, parent *argparse.Command, currGroup string, index int) {
	if index >= len(groups) {
		return
	}

	if currGroup == "" {
		currGroup += groups[index]
	} else {
		currGroup += " " + groups[index]
	}

	subParser, subParserExists := p.subParsers[currGroup]
	if subParserExists {
		p.createSubParserFor(groups, subParser, currGroup, index+1)
	} else {
		newSubparser := parent.NewCommand(groups[index], "")
		p.subParsers[currGroup] = newSubparser
		p.createSubParserFor(groups, newSubparser, currGroup, index+1)
	}
}

type Parser struct {
	argParser   *argparse.Parser
	subParsers  map[string]*argparse.Command
	name        string
	description string
}
