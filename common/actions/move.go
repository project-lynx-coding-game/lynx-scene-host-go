package actions

import (
	"encoding/json"

	"github.com/group-project-gut/lynx-scene-host/common"
)

type Move struct {
	ObjectId int64         `json:"object_id"`
	Vector   common.Vector `json:"vector"`
}

func (move Move) Type() string {
	return "Move"
}

func (move Move) Args() string {
	args, err := json.Marshal(move)
	if err != nil {
		panic(err)
	}
	return string(args)
}

func (move Move) Requirements(scene common.IScene) []func(scene common.IScene) bool {
	return make([]func(scene common.IScene) bool, 0, 0)
}

func (move Move) Effects(scene common.IScene) []func(scene common.IScene) common.IScene {
	return []func(scene common.IScene) common.IScene{
		func(scene common.IScene) common.IScene {
			object, err := scene.GetObjectById(move.ObjectId)

			if err != nil {
				return scene
			}

			// TODO: Maybe we should make Vector immutable or sth
			(*object).Position.Add(common.NORTH())
			return scene
		},
	}
}
