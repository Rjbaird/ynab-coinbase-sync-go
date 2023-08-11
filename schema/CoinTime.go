package schema

import "time"

type CoinTime struct {
	Data CoinTimeData `json:"data"`
}
type CoinTimeData struct {
	Iso   time.Time `json:"iso"`
	Epoch int     `json:"epoch"`
}