package common

type IAction interface {
	Requirements(scene IScene) []func(scene IScene) bool
	Effects(scene IScene) []func(scene IScene) IScene
}
