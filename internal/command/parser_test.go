package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSubParser_CreatesCommandTree(t *testing.T) {
	parser := NewParser("Demo", "Demo CLI")

	subParser := parser.getSubParser("mammals placentilia")

	assert.Equal(t, "mammals", subParser.GetParent().GetName())
	assert.Equal(t, "placentilia", subParser.GetName())
}
