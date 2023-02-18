package main

import (
	"encoding/json"
	"fmt"

	. "github.com/group-project-gut/lynx-scene-host/common"
	. "github.com/group-project-gut/lynx-scene-host/common/actions"
	. "github.com/group-project-gut/lynx-scene-host/common/objects"
	. "github.com/group-project-gut/lynx-scene-host/common/scenes"
)

func main() {
	simple_scene := NewSimpleScene()
	var scene Scene = &simple_scene
	scene.AddObject(&Floor{BaseObject: BaseObject{BaseId: 0, BasePosition: Vector{X: 0, Y: 0}}})
	agent := Agent{BaseObject: BaseObject{BaseId: 1, BasePosition: Vector{X: 0, Y: 0}}}
	agent.Initialize()
	agent.Tick(scene)
	scene.AddObject(&agent)

	move := Move{TargetId: 0, Vector: NORTH()}
	for _, effect := range move.Effects(scene) {
		scene = effect(scene)
	}

	sceneMarshalled, _ := json.Marshal(scene)
	fmt.Println(string(sceneMarshalled))
}
