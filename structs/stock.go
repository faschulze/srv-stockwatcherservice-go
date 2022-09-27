package structs

// Position struct
type Position struct {
	Name         string
	Tickersymbol string
	ReqURL       string
	CurrentPrice float32
	PositionCategory
	PositionLimit
}

// PositionLimit struct
type PositionLimit struct {
	LimitType  string
	LimitPrice float32
}

// PositionCategory struct
type PositionCategory struct {
	Category string
}
