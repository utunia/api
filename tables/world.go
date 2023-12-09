package tables

type World struct {
	ID          ID     `json:"id"`
	Name        string `json:"name"`
	NationCount uint64 `json:"nation_count"`
	Nations     []ID   `json:"nations"`
}
