package common

type IScene interface {
	AddObject(object *Object)
	GetObjectById(id int64) (*Object, error)
	UnmarshalJSON(data []byte) error
	MarshalJSON() ([]byte, error)
}
