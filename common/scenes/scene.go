package scenes

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"

	. "github.com/group-project-gut/lynx-scene-host/common"
	. "github.com/group-project-gut/lynx-scene-host/common/entity"
)

type Scene struct {
	Entities   []Entity `json:"entities"`
	idMap      map[int64]*Object
	processMap map[string]*process
}

type process struct {
	Encoder *json.Encoder
	Decoder *json.Decoder
}

type serializedEntity struct {
	Type string `json:"type"`
	Args string `json:"args"`
}

type exportedScene struct {
	Entities []serializedEntity `json:"entities"`
}

func (scene *Scene) AddObject(object *Object) {
	scene.Entities = append(scene.Entities, Entity{IEntity: object})

	scene.idMap[object.Id] = object
}

func (scene *Scene) CreateProcess(name string) {
	//cmd := exec.Command("python", "-m", "debugpy", "--wait-for-client", "--listen", "0.0.0.0:5678", "./execution-runtime.py")
	cmd := exec.Command("python", "./execution-runtime.py")

	child_read, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	child_write, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	cmd.Stderr = os.Stderr

	scene.processMap[name] = &process{Encoder: json.NewEncoder(child_write), Decoder: json.NewDecoder(child_read)}

	cmd.Start()
	if err != nil {
		panic(err)
	}
}

func (scene *Scene) RunObject(object *Object) []IAction {
	_, found := scene.processMap[object.Name]
	if !found {
		scene.CreateProcess(object.Name)
	}
	process := scene.processMap[object.Name]

	err := process.Encoder.Encode(scene)
	if err != nil {
		panic(err)
	}

	var new_scene Scene
	err = process.Decoder.Decode(&new_scene)
	if err != nil {
		panic(err)
	}

	for _, entity := range new_scene.Entities {
		action, ok := entity.IEntity.(IAction)
		if ok {
			for _, effect := range action.Effects(&new_scene) {
				scene = effect(scene).(*Scene)
			}
		}
	}

	return make([]IAction, 0, 0)
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
		Entities:   make([]Entity, 0, 32),
		idMap:      map[int64]*Object{},
		processMap: map[string]*process{},
	}
}
