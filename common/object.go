package common

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
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
	encoder             json.Encoder
	decoder             json.Decoder
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

func (object *Object) Initialize() {
	cmd := exec.Command("python", "-m", "debugpy", "--wait-for-client", "--listen", "0.0.0.0:5678", "./Object.py")
	//cmd := exec.Command("python", "./Object.py")

	child_read, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	child_write, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	cmd.Stderr = os.Stderr

	object.encoder = *json.NewEncoder(child_write)
	object.decoder = *json.NewDecoder(child_read)

	cmd.Start()
	if err != nil {
		panic(err)
	}
}

func (object *Object) runCode(scene IScene) []IAction {
	err := object.encoder.Encode(scene)
	if err != nil {
		panic(err)
	}

	var new_scene IScene
	err = object.decoder.Decode(&new_scene)
	if err != nil {
		panic(err)
	}

	sceneMarshalled, _ := json.Marshal(new_scene)
	fmt.Println(string(sceneMarshalled))

	return make([]IAction, 0, 0)
}
