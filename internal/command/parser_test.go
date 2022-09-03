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

func TestGetSubParser_ParentCanHaveChildren(t *testing.T) {
	parser := NewParser("Demo", "Demo CLI")

	cold := parser.getSubParser("mammals cold")
	warm := parser.getSubParser("mammals warm")

	assert.Equal(t, "mammals", cold.GetParent().GetName())
	assert.Equal(t, "mammals", warm.GetParent().GetName())

}
