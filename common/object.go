package common

type Object interface {
	Tick(scene Scene) []Action
	Position() *Vector
	Id() int64
}
