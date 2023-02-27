package common

import (
	"encoding/json"
	"testing"
)

func TestObjectType(t *testing.T) {
	object := Object{
		Id:                  123,
		Name:                "Ondrejek",
		Position:            Vector{1123, -1231},
		AdditionalPositions: []Vector{{0, 1}, {1, 1}},
		State:               "smoking",
		Walkable:            false,
		Tick:                "move(Direction.SOUTH)",
		OnDeath:             "wave()",
	}
	entity := object

	type_string := entity.Type()
	if type_string != "Object" {
		t.Fatalf(`Object.Type() error. Expected "Object", received "%s"`, type_string)
	}
}

// TODO: We should probably test this by the `Args()` interface
func TestObjectSerialization(t *testing.T) {
	object := Object{
		Id:                  123,
		Name:                "Ondrejek",
		Position:            Vector{1123, -1231},
		AdditionalPositions: []Vector{{0, 1}, {1, 1}},
		State:               "smoking",
		Walkable:            false,
		Tick:                "move(Direction.SOUTH)",
		OnDeath:             "wave()",
		Owner:               "rut",
	}

	object_serialized, err := json.Marshal(object)
	if err != nil {
		t.Fatalf(`Failed to serialize Object - %v`, err)
	}

	expected_object_serialized := `{"id":123,"name":"Ondrejek","position":{"x":1123,"y":-1231},"additional_positions":[{"x":0,"y":1},{"x":1,"y":1}],"state":"smoking","walkable":false,"tick":"move(Direction.SOUTH)","on_death":"wave()","owner":"rut"}`
	if string(object_serialized) != expected_object_serialized {
		t.Fatalf(`Serialization format is wrong - got: %s, expected: %s`, string(object_serialized), expected_object_serialized)
	}
}

func TestObjectDeserialization(t *testing.T) {
	expected_object := Object{
		Id:                  123,
		Name:                "Ondrejek",
		Position:            Vector{1123, -1231},
		AdditionalPositions: []Vector{{0, 1}, {1, 1}},
		State:               "smoking",
		Walkable:            false,
		Tick:                "move(Direction.SOUTH)",
		OnDeath:             "wave()",
	}

	object_serialized := `{"id":123,"name":"Ondrejek","position":{"x":1123,"y":-1231},"additional_positions":[{"x":0,"y":1},{"x":1,"y":1}],"state":"smoking","walkable":false,"tick":"move(Direction.SOUTH)","on_death":"wave()"}`

	object_deserialized := Object{}
	err := json.Unmarshal([]byte(object_serialized), &object_deserialized)
	if err != nil {
		t.Fatalf(`Failed to deserialize Object - %v`, err)
	}

	if expected_object.Name != object_deserialized.Name {
		t.Fatalf(`Deserialized struct is wrong`)
	}
}

func TestObjectPartialDeserialization(t *testing.T) {
	expected_object := Object{
		Id: 123,
	}

	object_serialized := `{"id":123}`

	object_deserialized := Object{}
	err := json.Unmarshal([]byte(object_serialized), &object_deserialized)
	if err != nil {
		t.Fatalf(`Failed to deserialize Object - %v`, err)
	}

	if expected_object.Id != object_deserialized.Id {
		t.Fatalf(`Deserialized struct is wrong`)
	}
	if expected_object.Name != object_deserialized.Name {
		t.Fatalf(`Deserialized struct is wrong`)
	}
}
