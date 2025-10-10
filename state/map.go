package state

type Map struct {
	MapID        int          `json:"map_id"`
	Name         string       `json:"name"`
	Skin         string       `json:"skin"`
	X            int          `json:"x"`
	Y            int          `json:"y"`
	Layer        string       `json:"layer"`
	Access       Access       `json:"access"`
	Interactions Interactions `json:"interactions"`
}

type Access struct {
	Type       string      `json:"type"`
	Conditions []Condition `json:"conditions"`
}

type Interactions struct {
	Content    *InteractionContent `json:"content,omitempty"`
	Transition *Transition         `json:"transition,omitempty"`
}

type InteractionContent struct {
	Type string `json:"type"`
	Code string `json:"code"`
}

type Transition struct {
	MapID      int         `json:"map_id"`
	X          int         `json:"x"`
	Y          int         `json:"y"`
	Layer      string      `json:"layer"`
	Conditions []Condition `json:"conditions"`
}

type Condition struct {
	Code     string      `json:"code"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}
