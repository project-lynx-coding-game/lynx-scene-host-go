package scenes

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	. "github.com/group-project-gut/lynx-scene-host/common"
	. "github.com/group-project-gut/lynx-scene-host/common/objects"
)

type SimpleScene struct {
	types   []string
	objects []Object
	idMap   map[int64]*Object
}

type exportedSimpleScene struct {
	Types   []string `json:"types"`
	Objects []string `json:"objects"`
}

// `types` and `objects` are private, so we do
// the little trick in order to have them exported
// into a `json`
func (scene SimpleScene) MarshalJSON() ([]byte, error) {
	exported_scene := exportedSimpleScene{
		Types:   scene.types,
		Objects: make([]string, len(scene.objects)),
	}

	for index, object := range scene.objects {
		object_json, err := json.Marshal(object)
		if err != nil {
			return nil, err
		}
		exported_scene.Objects[index] = string(object_json)
	}

	scene_json, err := json.Marshal(exported_scene)
	if err != nil {
		return nil, err
	}

	return scene_json, nil
}

func (scene *SimpleScene) UnmarshalJSON(data []byte) error {
	var exported_scene exportedSimpleScene
	err := json.Unmarshal(data, &exported_scene)
	if err != nil {
		return err
	}

	for index, object_str := range exported_scene.Objects {
		var object Object

		switch exported_scene.Types[index] {
		case "Floor":
			object = &Floor{}
		case "Agent":
			object = &Agent{}
		}

		err := json.Unmarshal([]byte(object_str), &object)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func (scene *SimpleScene) AddObject(object Object) {
	scene.objects = append(scene.objects, object)

	// Fetch `object` type by using `%T` format option
	objectNameParts := strings.Split(fmt.Sprintf("%T", object), ".")
	scene.types = append(scene.types, objectNameParts[len(objectNameParts)-1])
	scene.idMap[object.Id()] = &scene.objects[len(scene.objects)-1]
}

func (scene *SimpleScene) GetObjectById(id int64) (*Object, error) {
	object := scene.idMap[id]
	if object == nil {
		return nil, errors.New("Could not find an object with the Id!")
	}

	return object, nil
}

func NewSimpleScene() SimpleScene {
	return SimpleScene{
		types:   make([]string, 0, 32),
		objects: make([]Object, 0, 32),
		idMap:   map[int64]*Object{},
	}
}
