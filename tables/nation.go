package tables

type Nation struct {
	ID             ID     `json:"id"`
	Name           string `json:"name"`
	PassKey        string `json:"passkey"`
	Flag           string `json:"flag"`
	Motto          string `json:"motto"`
	Currency       string `json:"currency"`
	Classification string `json:"classification"`
	Ideology       string `json:"ideology"`
	Stats          Stats  `json:"stats"`
}

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

	Corruption    StatsPercentage `json:"corruption"`
	HealthCare    StatsPercentage `json:"healthcare"`
	Argiculture   StatsPercentage `json:"argiculture"`
	TerrorismRate StatsPercentage `json:"terrorism_rate"`
	Sprituality   StatsPercentage `json:"sprituality"`
	Education     StatsPercentage `json:"education"`
	SecularRate   StatsPercentage `json:"secular_rate"`
}
