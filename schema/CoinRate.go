package schema

import "time"

type CoinRate struct {
	AssetIDBase string  `json:"asset_id_base"`
	Rates       []Rates `json:"rates"`
}
type Rates struct {
	Time         time.Time `json:"time"`
	AssetIDQuote string    `json:"asset_id_quote"`
	Rate         float64   `json:"rate"`
}