package common

import (
	"encoding/json"
)

// Right now there's no constructor for the type
// because it is not intended for `scene-host` to
// be responsible for creation of the structure.
// That is - it should be created by `scene-generator`
type Object struct {
	Id                  int64    `json:"id"`
	Name                string   `json:"name"`
	Position            Vector   `json:"position"`
	AdditionalPositions []Vector `json:"additional_positions"`
	State               string   `json:"state"`
	Walkable            bool     `json:"walkable"`
	Tick                string   `json:"tick"`
	OnDeath             string   `json:"on_death"`
	Owner               string   `json:"owner"`
}

func (object Object) Type() string {
	return "Object"
}

func (object Object) Args() string {
	args, err := json.Marshal(object)
	if err != nil {
		// TODO: we might add returning of some log instead?
		panic(err)
	}
	return string(args)
}
