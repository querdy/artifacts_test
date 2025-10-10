package state

type DepositItem struct {
	Cooldown  Cooldown  `json:"cooldown"`
	Items     []Item    `json:"items"`
	Bank      []Item    `json:"bank"`
	Character Character `json:"character"`
}
