package tables

type World struct {
	Name        string
	ID          string
	NationCount uint64
	Nations     []Nation
}
