package actions

import (
	"encoding/json"
	"testing"

	"github.com/group-project-gut/lynx-scene-host/common"
)

func TestMoveType(t *testing.T) {
	move := Move{
		TargetId: 1245,
		Vector:   common.Vector{X: 321, Y: 565},
	}
	entity := move

	type_string := entity.Type()
	if type_string != "Move" {
		t.Fatalf(`Move.Type() error. Expected "Move", received "%s"`, type_string)
	}
}
func TestMoveArgs(t *testing.T) {
	move := Move{
		TargetId: 1245,
		Vector:   common.Vector{X: 321, Y: 565},
	}
	entity := move

	args := entity.Args()
	expected := `{"target_id":1245,"vector":{"x":321,"y":565}}`
	if args != expected {
		t.Fatalf(`Move.Args() error. Expected "%s", received "%s"`, expected, args)
	}
}

func TestMoveDeserialization(t *testing.T) {
	move := Move{
		TargetId: 1245,
		Vector:   common.Vector{X: 321, Y: 565},
	}
	args := `{"target_id":1245,"vector":{"x":321,"y":565}}`

	var move_deserialized Move
	err := json.Unmarshal([]byte(args), &move_deserialized)
	if err != nil {
		t.Fatalf(`Failed to deserialize Object - %v`, err)
	}

	if move != move_deserialized {
		t.Fatalf(`Deserialized object is not equal. Expected "%+v", received "%+v"`, move, move_deserialized)
	}
}

func TestMoveEffects(t *testing.T) {
	// TODO: test if applying effects of `Move` results in proper state
}

func TestMoveRequirements(t *testing.T) {
	// TODO: test if requirements of `Move` properly discard malformed actions
}
