package tables

type World struct {
	ID          ID
	Name        string
	NationCount uint64
	Nations     []ID
}
