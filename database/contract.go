package database

type Modeler interface {
	Find(Id int) (model Modeler, err error)
}
