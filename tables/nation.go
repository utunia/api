package tables

type Nation struct {
	Name           string
	PassKey        string
	Flag           string
	Motto          string
	Currency       string
	Classification string
	Ideology       string
	World          World
	Stats          Stats
}
