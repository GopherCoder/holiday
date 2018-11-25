package model

type BasicSerialization struct {
	ID       uint   `json:"id"`
	StartDay string `json:"start_day"`
	EndDay   string `json:"end_day"`
	Count    int    `json:"count"`
	EnName   string `json:"en_name"`
	ChName   string `json:"ch_name"`
}
