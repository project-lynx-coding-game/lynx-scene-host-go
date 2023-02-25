package scenes

import (
	"encoding/json"
	"errors"

	. "github.com/group-project-gut/lynx-scene-host/common"
	. "github.com/group-project-gut/lynx-scene-host/common/actions"
)

type Scene struct {
	entities []IEntity
	idMap    map[int64]*Object
}

type serializedEntity struct {
	Type string `json:"type"`
	Args string `json:"args"`
}

type exportedScene struct {
	Entities []serializedEntity `json:"entities"`
}

// `types` and `objects` are private, so we do
// the little trick in order to have them exported
// into a `json`
func (scene Scene) MarshalJSON() ([]byte, error) {
	exported_scene := exportedScene{
		Entities: make([]serializedEntity, len(scene.entities)),
	}

	for index, entity := range scene.entities {
		exported_scene.Entities[index] = serializedEntity{entity.Type(), entity.Args()}
	}

	scene_json, err := json.Marshal(exported_scene)
	if err != nil {
		return nil, err
	}

	return scene_json, nil
}

func (scene *Scene) UnmarshalJSON(data []byte) error {
	var exported_scene exportedScene
	err := json.Unmarshal(data, &exported_scene)
	if err != nil {
		return err
	}

	for _, serialized_entity := range exported_scene.Entities {
		var entity IEntity

		// Here we `MUST` put all the structs that we would like to
		// deserialize to.
		switch serialized_entity.Type {
		case "Object":
			var object Object
			err := json.Unmarshal([]byte(serialized_entity.Args), &object)
			if err != nil {
				panic(err)
			}
			entity = object
		case "Move":
			var move Move
			err := json.Unmarshal([]byte(serialized_entity.Args), &move)
			if err != nil {
				panic(err)
			}
			entity = move
		}
		scene.entities = append(scene.entities, entity)
	}

	return nil
}

func (scene *Scene) AddObject(object *Object) {
	scene.entities = append(scene.entities, object)

	scene.idMap[object.Id] = object
}

func (scene *Scene) GetObjectById(id int64) (*Object, error) {
	object := scene.idMap[id]
	if object == nil {
		return nil, errors.New("Could not find an object with the Id!")
	}

	return object, nil
}

func NewScene() Scene {
	return Scene{
		entities: make([]IEntity, 0, 32),
		idMap:    map[int64]*Object{},
	}
}
