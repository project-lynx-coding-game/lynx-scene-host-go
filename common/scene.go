package common

type Scene interface {
	AddObject(object Object)
	GetObjectById(id int64) (*Object, error)
	UnmarshalJSON(data []byte) error
}
