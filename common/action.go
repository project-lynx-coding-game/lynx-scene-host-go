package common

type Action interface {
	Requirements(scene Scene) []func(scene Scene) bool
	Effects(scene Scene) []func(scene Scene) Scene
}
