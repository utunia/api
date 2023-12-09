package tables

type Nation struct {
	ID             ID
	Name           string
	PassKey        string
	Flag           string
	Motto          string
	Currency       string
	Classification string
	Ideology       string
	Stats          Stats
}

type NeutralResource struct {
	Name       string
	Percentage StatsPercentage
}

type Export struct {
	Name       string
	Percentage StatsPercentage
}

type Stats struct {
	Population       StatsNumber
	NeutralResources []NeutralResource
	GDP              StatsNumber
	Exports          []Export

	Power     StatsNumber
	Army      StatsNumber
	Navy      StatsNumber
	AirForces StatsNumber
	Weapons   StatsNumber
	Nukes     StatsNumber

	Corruption    StatsPercentage
	HealthCare    StatsPercentage
	Argiculture   StatsPercentage
	TerrorismRate StatsPercentage
	Sprituality   StatsPercentage
	Education     StatsPercentage
	SecularRate   StatsPercentage
}
