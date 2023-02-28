package entity

import (
	"encoding/json"

	. "github.com/group-project-gut/lynx-scene-host/common"
	. "github.com/group-project-gut/lynx-scene-host/common/actions"
)

type Entity struct {
	IEntity
}

type exportedEntity struct {
	Type string `json:"type"`
	Args string `json:"args"`
}

func (entity Entity) MarshalJSON() ([]byte, error) {
	exported_entity := exportedEntity{entity.Type(), entity.Args()}
	serialized_entity, err := json.Marshal(exported_entity)
	if err != nil {
		return nil, err
	}
	return serialized_entity, nil
}

func (entity *Entity) UnmarshalJSON(data []byte) error {
	var exported_entity exportedEntity
	err := json.Unmarshal(data, &exported_entity)
	if err != nil {
		return nil
	}

	// Here we `MUST` put all the structs that we would like to
	// deserialize to.
	switch exported_entity.Type {
	case "Object":
		var object Object
		err := json.Unmarshal([]byte(exported_entity.Args), &object)
		if err != nil {
			panic(err)
		}
		entity.IEntity = object
	case "Move":
		var move Move
		err := json.Unmarshal([]byte(exported_entity.Args), &move)
		if err != nil {
			panic(err)
		}
		entity.IEntity = move
	}

	return nil
}
