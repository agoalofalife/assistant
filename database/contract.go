package database

type Database interface {
	Open(...interface{}) (interface{}, error)
}

type Model interface {
	Find(Id int) (stringJson []byte, err error)
}
