package scenes

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/group-project-gut/lynx-scene-host/common"
)

func TestSceneSerialization(t *testing.T) {
	object := common.Object{Id: 213123}
	scene := NewScene()
	scene.AddObject(&object)

	scene_serialized, err := json.Marshal(scene)
	if err != nil {
		t.Fatalf(`Failed to serialize Object - %v`, err)
	}

	// Double quotes (`"`) have to be escaped properly
	expected := fmt.Sprintf(`{"entities":[{"type":"Object","args":"%s"}]}`, strings.ReplaceAll(object.Args(), `"`, `\"`))
	if string(scene_serialized) != expected {
		t.Fatalf(`Move.Args() error. Expected "%s", received "%s"`, expected, scene_serialized)
	}
}
