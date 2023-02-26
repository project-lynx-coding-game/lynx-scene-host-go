package main

import (
	"encoding/json"
	"fmt"

	. "github.com/group-project-gut/lynx-scene-host/common"
	_ "github.com/group-project-gut/lynx-scene-host/common/actions"
	. "github.com/group-project-gut/lynx-scene-host/common/scenes"
)

func main() {
	simple_scene := NewScene()
	var scene IScene = &simple_scene
	agent := Object{Id: 123, AdditionalPositions: make([]Vector, 0), Tick: "print(42)"}
	scene.AddObject(&agent)
	scene.RunObject(&agent)

	// move := Move{TargetId: 0, Vector: NORTH()}
	// for _, effect := range move.Effects(scene) {
	// 	scene = effect(scene)
	// }

	sceneMarshalled, _ := json.Marshal(scene)
	fmt.Println(string(sceneMarshalled))
}
