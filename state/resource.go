package state

type Resource struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Skill string `json:"skill"`
	Level int    `json:"level"`
	Drops []Drop `json:"drops"`
}

type Drop struct {
	Code        string `json:"code"`
	Rate        int    `json:"rate"`
	MinQuantity int    `json:"min_quantity"`
	MaxQuantity int    `json:"max_quantity"`
}
