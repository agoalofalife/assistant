package database

type Collector interface {
	ToJson() ([]byte, error)
}
