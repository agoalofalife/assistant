package database


//import "github.com/agoalofalife/assistant"

//type Modeler interface {
//	Find(Id int) (model Modeler, err error)
//	Collector
//}

type Collector interface {
	ToJson() ([]byte, error)
}
