package tables

type StatsNumber int32
type StatsPercentage float32

type NeutralResource struct {
    Name string
    Percentage StatsPercentage
}

type Stats struct {
    Population StatsNumber
    NeutralResources []NeutralResource
    GDP StatsNumber
    Exports StatsPercentage

    Power StatsNumber
    Army StatsNumber
    Navy StatsNumber
    AirForces StatsNumber
    Weapons StatsNumber
    Nukes StatsNumber

    Corruption StatsPercentage
    HealthCare StatsPercentage
    Argiculture StatsPercentage
    TerrorismRate StatsPercentage
    Sprituality StatsPercentage
    Education StatsPercentage
    SecularRate StatsPercentage
}
