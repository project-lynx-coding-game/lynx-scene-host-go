package objects

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	. "github.com/group-project-gut/lynx-scene-host/common"
)

type Agent struct {
	BaseObject
	encoder json.Encoder
	decoder json.Decoder
}

func (agent *Agent) Initialize() {
	cmd := exec.Command("python", "-m", "debugpy", "--wait-for-client", "--listen", "0.0.0.0:5678", "./agent.py")
	//cmd := exec.Command("python", "./agent.py")

	child_read, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	child_write, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	cmd.Stderr = os.Stderr

	agent.encoder = *json.NewEncoder(child_write)
	agent.decoder = *json.NewDecoder(child_read)

	cmd.Start()
	if err != nil {
		panic(err)
	}
}

func (agent *Agent) Tick(scene Scene) []Action {
	err := agent.encoder.Encode(scene)
	if err != nil {
		panic(err)
	}

	var new_scene Scene
	err = agent.decoder.Decode(&new_scene)
	if err != nil {
		panic(err)
	}

	sceneMarshalled, _ := json.Marshal(new_scene)
	fmt.Println(string(sceneMarshalled))

	return make([]Action, 0, 0)
}
