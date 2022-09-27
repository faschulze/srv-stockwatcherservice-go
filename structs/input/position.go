package input

// InputStockPosition struct
type InputStockPosition struct {
	Name              string  `csv:"name"`
	Tickersymbol      string  `csv:"tickersymbol"`
	PositionLimitType string  `csv:"limit_type"`
	NotionID          string  `csv:"notion_id"`
	ISIN              string  `csv:"isin"`
	PositionLimit     float32 `csv:"position_limit"`
}

// InputCryptoPosition struct
type InputCryptoPosition struct {
	Name              string  `csv:"name"`
	Tickersymbol      string  `csv:"tickersymbol"`
	PositionLimitType string  `csv:"limit_type"`
	NotionID          string  `csv:"notion_id"`
	PositionLimit     float32 `csv:"position_limit"`
}

// InputFondsPosition struct
type InputFondsPosition struct {
	Name              string  `csv:"name"`
	Tickersymbol      string  `csv:"tickersymbol"`
	PositionLimitType string  `csv:"limit_type"`
	NotionID          string  `csv:"notion_id"`
	ISIN              string  `csv:"isin"`
	PositionLimit     float32 `csv:"position_limit"`
}

// InputIndexPosition struct
type InputIndexPosition struct {
	Name              string  `csv:"name"`
	Tickersymbol      string  `csv:"tickersymbol"`
	PositionLimitType string  `csv:"limit_type"`
	ISIN              string  `csv:"isin"`
	PositionLimit     float32 `csv:"position_limit"`
}
