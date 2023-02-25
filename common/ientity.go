package common

type IEntity interface {
	Type() string
	Args() string
}
