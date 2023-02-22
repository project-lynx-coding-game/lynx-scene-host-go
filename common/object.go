package common

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
}
