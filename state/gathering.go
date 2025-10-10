package state

type Gathering struct {
	Cooldown  Cooldown  `json:"cooldown"`
	Details   Details   `json:"details"`
	Character Character `json:"character"`
}

type Details struct {
	XP    int    `json:"xp"`
	Items []Item `json:"items"`
}

type Item struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}
