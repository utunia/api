package tables

type World struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	NationCount uint64   `json:"nation_count"`
	Nations     []string `json:"nations"`
}
