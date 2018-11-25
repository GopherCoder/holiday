package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Holiday struct {
	gorm.Model
	EnName     string `gorm:"type:varchar(20)" json:"en_name"`
	ChName     string `gorm:"type:varchar(20)" json:"ch_name"`
	StartYear  int    `gorm:"type:integer" json:"start_year"`
	EndYear    int    `gorm:"type:integer" json:"end_year"`
	StartMonth int    `gorm:"type:integer" json:"start_month"`
	EndMonth   int    `gorm:"type:integer" json:"end_month"`
	StartDay   int    `gorm:"type:integer" json:"start_day"`
	EndDay     int    `gorm:"type:integer" json:"end_day"`
}

func (h *Holiday) TableName() string {
	return "tb_holidays"
}

func (h Holiday) handlerString() (string, string, float64) {
	startDay := time.Date(h.StartYear, time.Month(h.StartMonth), h.StartDay, 0, 0, 0, 0, time.Local)
	endDay := time.Date(h.EndYear, time.Month(h.EndMonth), h.EndDay, 0, 0, 0, 0, time.Local)
	count := endDay.Sub(startDay).Hours() / 24
	return startDay.Format("2006/01/02"), endDay.Format("2006/01/02"), count + 1
}

func (h Holiday) Basic() BasicSerialization {
	s, e, c := h.handlerString()
	return BasicSerialization{
		ID:       h.ID,
		EnName:   h.EnName,
		ChName:   h.ChName,
		StartDay: s,
		EndDay:   e,
		Count:    int(c),
	}
}
