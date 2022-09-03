# commando
An elegant CLI framework 

## Getting Started

### Installation

To install commando simply do

`go get github.com/jobala/commando`

### Usage

Create a CLI instance

```go
cli := commando.NewCli("animals", "A cli app about the animal kingdom")

```

Commando exposes a declarative api for grouping commands and adding commands to a command group. The code snippet below adds the following commands;

1. `$animals vertebrates warm mammals bear --loudnes <int>`
2. `$animals vertebrates warm mammals cow --loudness <int>`


You can use **--help** to show the available commands in a command group.
`$animals --help` will list available sub command[group]s

```go 
cli.NewCommandGroup("vertebrates warm mammals").
    WithCommand(commando.Command("bear", Bear)).
    WithCommand(commando.Command("cow", Cow)).

func Bear(args BearArgs) {
    // do something with BearArgs and print result
}

type BearArgs struct {
	Loudness int
}

func Cow(args CowArgs) {
    // do something CowArgs and print result
}

type CowArgs struct {
	Loudness int
}
```

The handler function takes a single struct as an argument and the struct fields will be used to create the command's flags. The fields **must** be public.


To add another command group under `vertebrates`, do the following

```go
cli.NewCommandGroup("vertebrates cold").
    WithCommand(commando.Command("snake", Snake))

```
Now you'll have `$animals vertebrates cold ...` on top of the initially defined commands

To add another command group under animals, simply do the following

```go
cli.NewCommandGroup("invertebrates").
    WithCommand(commando.Command("leech", handlerFunc)).

```

The snippet below will add `$animals invertebrates leech`
