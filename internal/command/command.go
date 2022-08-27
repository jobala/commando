package command

import (
	"encoding/json"
)

func (c Command[T]) Execute() {
	var arg T

	jsonStr, _ := json.Marshal(c.args)
	json.Unmarshal([]byte(jsonStr), &arg)

	c.handlerFunc(arg)
}

type Executor interface {
	Execute()
}

type Command[T any] struct {
	args        map[string]any
	handlerFunc func(args T)
}

type CommandTable = map[string][]Loader
