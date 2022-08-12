package main

import (
	"fmt"
	"os"

	"github.com/jobala/commando/pkg/command"
)

/**
TODO
    1. Move to a Go workspace
    2. Make project a Go package
*/
func main() {
	command.NewCommandGroup("animals", "animals you may know").
		WithCommand(command.NewCommand("cat", Meow)).
		WithCommand(command.NewCommand("cow", Cow))

	err := command.Parser.Parse(os.Args)
	if err != nil {
		fmt.Println(err)
	}

	for _, cmdItem := range command.CommandTable {
		if cmdItem.Cmd.Happened() {
			cmdItem.Executor.Execute()
		}
	}

}

func Meow(args CatArgs) {
	res := "meow"

	for i := 0; i < args.Loudness; i++ {
		res += "meow"
	}

	fmt.Println(res)
}

type CatArgs struct {
	Loudness int
}

func Cow(args CowArgs) {
	res := "moo"

	for i := 0; i < args.Loudness; i++ {
		res += "moo"
	}

	fmt.Println(res)
}

type CowArgs struct {
	Loudness int
}
