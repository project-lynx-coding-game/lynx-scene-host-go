package objects

import (
	. "github.com/group-project-gut/lynx-scene-host/common"
)

type BaseObject struct {
	BaseId       int64  `json:"id"`
	BasePosition Vector `json:"position"`
}

func (object *BaseObject) Tick(scene Scene) []Action {
	return make([]Action, 0, 0)
}

func (object *BaseObject) Position() *Vector {
	return &object.BasePosition
}

func (object *BaseObject) Id() int64 {
	return object.BaseId
}
