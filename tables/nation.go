package tables

type Nation struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	PassKey        string `json:"passkey"`
	Flag           string `json:"flag"`
	Motto          string `json:"motto"`
	Currency       string `json:"currency"`
	Classification string `json:"classification"`
	Ideology       string `json:"ideology"`
	Stats          Stats  `json:"stats"`
}

type StatsNumber int32
type StatsPercentage float32

type NeutralResource struct {
	Name       string          `json:"name"`
	Percentage StatsPercentage `json:"percentage"`
}

type Export struct {
	Name       string          `json:"name"`
	Percentage StatsPercentage `json:"percentage"`
}

type Stats struct {
	Population       StatsNumber       `json:"population"`
	NeutralResources []NeutralResource `json:"neutral_resources"`
	GDP              StatsNumber       `json:"gdb"`
	Exports          []Export          `json:"exports"`

	Power     StatsNumber `json:"power"`
	Army      StatsNumber `json:"army"`
	Navy      StatsNumber `json:"navy"`
	AirForces StatsNumber `json:"air_forces"`
	Weapons   StatsNumber `json:"weapons"`
	Nukes     StatsNumber `json:"nukes"`
	Security  StatsNumber `json:"security"`

	CivilRights   StatsPercentage `json:"civil_rights"`
	Corruption    StatsPercentage `json:"corruption"`
	HealthCare    StatsPercentage `json:"healthcare"`
	Argiculture   StatsPercentage `json:"argiculture"`
	TerrorismRate StatsPercentage `json:"terrorism_rate"`
	Sprituality   StatsPercentage `json:"sprituality"`
	Education     StatsPercentage `json:"education"`
	SecularRate   StatsPercentage `json:"secular_rate"`
	DeathRate     StatsPercentage `json:"death_rate"`
}

func (store SupabaseStore) GetNationByName(nation *Nation, name string) error {
	return store.client.DB.From("nation").
		Select("*").
		Limit(1).
		Single().
		Eq("name", name).
		Execute(nation)
}

func (store SupabaseStore) UpdateNation(newNation Nation, id string) error {
	return store.client.DB.From("nation").Update(newNation).Eq("id", id).Execute(nil)
}

func (store SupabaseStore) InsertNation(nation Nation) error {
	return store.client.DB.From("nation").Insert(nation).Execute(nil)
}

func (store SupabaseStore) DeleteNation(id string) error {
	return store.client.DB.From("nation").Delete().Eq("id", id).Execute(nil)
}
