package state

import "time"

type CharacterMovement struct {
	Cooldown    Cooldown    `json:"cooldown"`
	Destination Destination `json:"destination"`
	Path        [][]int     `json:"path"`
	Character   Character   `json:"character"`
}

type Cooldown struct {
	TotalSeconds     int       `json:"total_seconds"`
	RemainingSeconds int       `json:"remaining_seconds"`
	StartedAt        time.Time `json:"started_at"`
	Expiration       time.Time `json:"expiration"`
	Reason           string    `json:"reason"`
}

type Destination struct {
	MapID        int          `json:"map_id"`
	Name         string       `json:"name"`
	Skin         string       `json:"skin"`
	X            int          `json:"x"`
	Y            int          `json:"y"`
	Layer        string       `json:"layer"`
	Access       Access       `json:"access"`
	Interactions Interactions `json:"interactions"`
}
