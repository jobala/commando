package command

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand_HandlerFunctionIsExecuted(t *testing.T) {
	parsedArgs := map[string]any{
		"fname": "John",
		"lname": "Doe",
	}

	cmd := Command[PrintArgs]{
		handlerFunc: printHandler,
		args:        parsedArgs,
	}

	cmd.Execute()

	assert.Equal(t, name, "John Doe")
}

var name string

type PrintArgs struct {
	Fname string
	Lname string
}

func printHandler(args PrintArgs) {
	name = fmt.Sprintf("%s %s", args.Fname, args.Lname)
}
