package state

type InventoryItem struct {
	Slot     int    `json:"slot"`
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}
