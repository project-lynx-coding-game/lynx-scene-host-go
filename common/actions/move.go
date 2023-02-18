package actions

import (
	"github.com/group-project-gut/lynx-scene-host/common"
)

type Move struct {
	TargetId int64
	Vector   common.Vector
}

func (move Move) Requirements(scene common.Scene) []func(scene common.Scene) bool {
	return make([]func(scene common.Scene) bool, 0, 0)
}

func (move Move) Effects(scene common.Scene) []func(scene common.Scene) common.Scene {
	return []func(scene common.Scene) common.Scene{
		func(scene common.Scene) common.Scene {
			object, err := scene.GetObjectById(move.TargetId)

			if err != nil {
				return scene
			}

			(*object).Position().Add(common.NORTH())
			return scene
		},
	}
}
