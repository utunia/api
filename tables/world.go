package tables

type World struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	NationCount uint64   `json:"nation_count"`
	Nations     []string `json:"nations"`
}

func (store SupabaseStore) GetWorldById(world *World, id string) error {
	return store.client.DB.From("world").
		Select("*").
		Limit(1).
		Single().
		Eq("id", id).
		Execute(world)
}

func (store SupabaseStore) GetWorldByNationId(world *World, id string) error {
	return store.client.DB.From("world").
		Select("*").
		Limit(1).
		Single().
		Cs("nations", []string{id}).
		Execute(world)
}

func (store SupabaseStore) UpdateWorld(newWorld World, id string) error {
	return store.client.DB.From("world").Update(newWorld).Eq("id", id).Execute(nil)
}
