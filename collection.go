package assistant

type Collector interface {
	ToJson() ([]byte, error)
}
