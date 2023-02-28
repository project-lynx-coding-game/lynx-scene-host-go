package entity

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/group-project-gut/lynx-scene-host/common"
)

func TestEntitySerialization(t *testing.T) {
	object := common.Object{Name: "Ondrejek"}
	entity := Entity{&object}

	expected := fmt.Sprintf(`{"type":"Object","args":"%s"}`, strings.ReplaceAll(object.Args(), `"`, `\"`))
	serialized_entity, err := json.Marshal(entity)
	if err != nil {
		t.Fatalf(`Failed to serialize entity - %v`, err)
	}

	if expected != string(serialized_entity) {
		t.Fatalf(`Serialization format is wrong - got: %s, expected: %s`, string(serialized_entity), expected)
	}
}

func TestEntityDeserialization(t *testing.T) {
	//TODO:
}
