package client

import (
	"artifacts/state"
)

type ServerStatusData struct {
	Data state.ServerStatus `json:"data"`
}

type AccountDetails struct {
	Data state.Account `json:"data"`
}

type CharactersData struct {
	Data []state.Character `json:"data"`
}

type MapData struct {
	Data  []state.Map `json:"data"`
	Total int         `json:"total"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Pages int         `json:"pages"`
}

type ResourceData struct {
	Data  []state.Resource `json:"data"`
	Total int              `json:"total"`
	Page  int              `json:"page"`
	Size  int              `json:"size"`
	Pages int              `json:"pages"`
}

type CharacterMovementData struct {
	Data state.CharacterMovement `json:"data"`
}

type GatheringData struct {
	Data state.Gathering `json:"data"`
}

type DepositItemData struct {
	Data state.DepositItem `json:"data"`
}
