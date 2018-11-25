package data

const (
	NewYearDay = iota
	SpringFestivalDay
	TombSweepingDay
	LaborDay
	DragonBoatFestivalDay
	NationalDay
	MidAutumnFestivalDay
)

var chHolidays = [...]string{
	"元旦",
	"春节",
	"清明节",
	"劳动节",
	"端午节",
	"中秋节",
	"国庆节",
}
var enHolidays = [...]string{
	"New Year\\'s Day",
	"Spring Festival",
	"Tomb-sweeping Day",
	"Labour Day",
	"Dragon Boat Festival",
	"Mid-autumn Festival",
	"National Day",
}

type HistoryHoliday struct {
	Data [][]YearCollection `json:"data"`
}

type YearCollection struct {
	Start  string `json:"start"`
	End    string `json:"end"`
	ChName string `json:"ch_name"`
	EnName string `json:"en_name"`
}

func (h *HistoryHoliday) Add(one []YearCollection) {
	h.Data = append(h.Data, one)
}

func NewHistoryHoliday() *HistoryHoliday {
	return &HistoryHoliday{
		Data: [][]YearCollection{
			holiday2019,
			holiday2018,
			holiday2017,
			holiday2016,
			//holiday2015,
			//holiday2014,
			//holiday2013,
			//holiday2012,
			//holiday2011,
			//holiday2010,
		},
	}
}

var holiday2019 = []YearCollection{
	{
		Start:  "2018/12/30",
		End:    "2019/01/01",
		ChName: chHolidays[NewYearDay],
		EnName: enHolidays[NewYearDay],
	},
	{
		Start:  "2019/02/04",
		End:    "2019/02/10",
		ChName: chHolidays[SpringFestivalDay],
		EnName: enHolidays[SpringFestivalDay],
	},
	{
		Start:  "2019/04/05",
		End:    "2019/04/07",
		ChName: chHolidays[TombSweepingDay],
		EnName: enHolidays[TombSweepingDay],
	},
	{
		Start:  "2019/04/29",
		End:    "2019/05/01",
		ChName: chHolidays[LaborDay],
		EnName: enHolidays[LaborDay],
	},
	{
		Start:  "2019/06/07",
		End:    "2019/06/09",
		ChName: chHolidays[DragonBoatFestivalDay],
		EnName: enHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2019/09/13",
		End:    "2019/09/15",
		ChName: chHolidays[MidAutumnFestivalDay],
		EnName: enHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2019/10/01",
		End:    "2019/10/07",
		ChName: chHolidays[NationalDay],
		EnName: enHolidays[NationalDay],
	},
}
var holiday2018 = []YearCollection{
	{
		Start:  "2017/12/30",
		End:    "2018/01/01",
		ChName: chHolidays[NewYearDay],
		EnName: enHolidays[NewYearDay],
	},
	{
		Start:  "2018/02/15",
		End:    "2018/02/21",
		ChName: chHolidays[SpringFestivalDay],
		EnName: enHolidays[SpringFestivalDay],
	},
	{
		Start:  "2018/04/05",
		End:    "2018/04/07",
		ChName: chHolidays[TombSweepingDay],
		EnName: enHolidays[TombSweepingDay],
	},
	{
		Start:  "2018/04/29",
		End:    "2018/05/01",
		ChName: chHolidays[LaborDay],
		EnName: enHolidays[LaborDay],
	},
	{
		Start:  "2018/06/16",
		End:    "2018/06/18",
		ChName: chHolidays[DragonBoatFestivalDay],
		EnName: enHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2018/09/22",
		End:    "2018/09/24",
		ChName: chHolidays[MidAutumnFestivalDay],
		EnName: enHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2018/10/01",
		End:    "2018/10/07",
		ChName: chHolidays[NationalDay],
		EnName: enHolidays[NationalDay],
	},
}
var holiday2017 = []YearCollection{
	{
		Start:  "2016/12/31",
		End:    "2017/01/02",
		ChName: chHolidays[NewYearDay],
		EnName: enHolidays[NewYearDay],
	},
	{
		Start:  "2017/01/27",
		End:    "2017/02/02",
		ChName: chHolidays[SpringFestivalDay],
		EnName: enHolidays[SpringFestivalDay],
	},
	{
		Start:  "2017/04/02",
		End:    "2017/04/04",
		ChName: chHolidays[TombSweepingDay],
		EnName: enHolidays[TombSweepingDay],
	},
	{
		Start:  "2017/04/29",
		End:    "2017/05/01",
		ChName: chHolidays[LaborDay],
		EnName: enHolidays[LaborDay],
	},
	{
		Start:  "2017/05/28",
		End:    "2017/05/30",
		ChName: chHolidays[DragonBoatFestivalDay],
		EnName: enHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2017/10/01",
		End:    "2017/10/08",
		ChName: chHolidays[NationalDay],
		EnName: enHolidays[NationalDay],
	},
}
var holiday2016 = []YearCollection{
	{
		Start:  "2016/01/01",
		End:    "2016/01/03",
		ChName: chHolidays[NewYearDay],
		EnName: enHolidays[NewYearDay],
	},
	{
		Start:  "2016/02/07",
		End:    "2016/02/13",
		ChName: chHolidays[SpringFestivalDay],
		EnName: enHolidays[SpringFestivalDay],
	},
	{
		Start:  "2016/04/02",
		End:    "2016/04/04",
		ChName: chHolidays[TombSweepingDay],
		EnName: enHolidays[TombSweepingDay],
	},
	{
		Start:  "2016/04/30",
		End:    "2016/05/02",
		ChName: chHolidays[LaborDay],
		EnName: enHolidays[LaborDay],
	},
	{
		Start:  "2016/06/09",
		End:    "2016/06/11",
		ChName: chHolidays[DragonBoatFestivalDay],
		EnName: enHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2016/09/15",
		End:    "2016/09/17",
		ChName: chHolidays[MidAutumnFestivalDay],
		EnName: enHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2016/10/01",
		End:    "2016/10/07",
		ChName: chHolidays[NationalDay],
		EnName: enHolidays[NationalDay],
	},
}
var holiday2015 = []YearCollection{
	{
		Start:  "2018/12/30",
		End:    "2019/01/01",
		ChName: chHolidays[NewYearDay],
		EnName: enHolidays[NewYearDay],
	},
	{
		Start:  "2019/02/04",
		End:    "2019/02/10",
		ChName: chHolidays[SpringFestivalDay],
		EnName: enHolidays[SpringFestivalDay],
	},
	{
		Start:  "2019/04/05",
		End:    "2019/04/07",
		ChName: chHolidays[TombSweepingDay],
		EnName: enHolidays[TombSweepingDay],
	},
	{
		Start:  "2019/04/29",
		End:    "2019/05/01",
		ChName: chHolidays[LaborDay],
		EnName: enHolidays[LaborDay],
	},
	{
		Start:  "2019/06/07",
		End:    "2019/06/09",
		ChName: chHolidays[DragonBoatFestivalDay],
		EnName: enHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2019/09/13",
		End:    "2019/09/15",
		ChName: chHolidays[MidAutumnFestivalDay],
		EnName: enHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2019/10/01",
		End:    "2019/10/07",
		ChName: chHolidays[NationalDay],
		EnName: enHolidays[NationalDay],
	},
}
var holiday2014 = []YearCollection{
	{
		Start:  "2018/12/30",
		End:    "2019/01/01",
		ChName: chHolidays[NewYearDay],
		EnName: enHolidays[NewYearDay],
	},
	{
		Start:  "2019/02/04",
		End:    "2019/02/10",
		ChName: chHolidays[SpringFestivalDay],
		EnName: enHolidays[SpringFestivalDay],
	},
	{
		Start:  "2019/04/05",
		End:    "2019/04/07",
		ChName: chHolidays[TombSweepingDay],
		EnName: enHolidays[TombSweepingDay],
	},
	{
		Start:  "2019/04/29",
		End:    "2019/05/01",
		ChName: chHolidays[LaborDay],
		EnName: enHolidays[LaborDay],
	},
	{
		Start:  "2019/06/07",
		End:    "2019/06/09",
		ChName: chHolidays[DragonBoatFestivalDay],
		EnName: enHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2019/09/13",
		End:    "2019/09/15",
		ChName: chHolidays[MidAutumnFestivalDay],
		EnName: enHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2019/10/01",
		End:    "2019/10/07",
		ChName: chHolidays[NationalDay],
		EnName: enHolidays[NationalDay],
	},
}
var holiday2013 = []YearCollection{
	{
		Start:  "2018/12/30",
		End:    "2019/01/01",
		ChName: chHolidays[NewYearDay],
		EnName: enHolidays[NewYearDay],
	},
	{
		Start:  "2019/02/04",
		End:    "2019/02/10",
		ChName: chHolidays[SpringFestivalDay],
		EnName: enHolidays[SpringFestivalDay],
	},
	{
		Start:  "2019/04/05",
		End:    "2019/04/07",
		ChName: chHolidays[TombSweepingDay],
		EnName: enHolidays[TombSweepingDay],
	},
	{
		Start:  "2019/04/29",
		End:    "2019/05/01",
		ChName: chHolidays[LaborDay],
		EnName: enHolidays[LaborDay],
	},
	{
		Start:  "2019/06/07",
		End:    "2019/06/09",
		ChName: chHolidays[DragonBoatFestivalDay],
		EnName: enHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2019/09/13",
		End:    "2019/09/15",
		ChName: chHolidays[MidAutumnFestivalDay],
		EnName: enHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2019/10/01",
		End:    "2019/10/07",
		ChName: chHolidays[NationalDay],
		EnName: enHolidays[NationalDay],
	},
}
var holiday2012 = []YearCollection{
	{
		Start:  "2018/12/30",
		End:    "2019/01/01",
		ChName: chHolidays[NewYearDay],
		EnName: enHolidays[NewYearDay],
	},
	{
		Start:  "2019/02/04",
		End:    "2019/02/10",
		ChName: chHolidays[SpringFestivalDay],
		EnName: enHolidays[SpringFestivalDay],
	},
	{
		Start:  "2019/04/05",
		End:    "2019/04/07",
		ChName: chHolidays[TombSweepingDay],
		EnName: enHolidays[TombSweepingDay],
	},
	{
		Start:  "2019/04/29",
		End:    "2019/05/01",
		ChName: chHolidays[LaborDay],
		EnName: enHolidays[LaborDay],
	},
	{
		Start:  "2019/06/07",
		End:    "2019/06/09",
		ChName: chHolidays[DragonBoatFestivalDay],
		EnName: enHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2019/09/13",
		End:    "2019/09/15",
		ChName: chHolidays[MidAutumnFestivalDay],
		EnName: enHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2019/10/01",
		End:    "2019/10/07",
		ChName: chHolidays[NationalDay],
		EnName: enHolidays[NationalDay],
	},
}
var holiday2011 = []YearCollection{
	{
		Start:  "2018/12/30",
		End:    "2019/01/01",
		ChName: chHolidays[NewYearDay],
		EnName: enHolidays[NewYearDay],
	},
	{
		Start:  "2019/02/04",
		End:    "2019/02/10",
		ChName: chHolidays[SpringFestivalDay],
		EnName: enHolidays[SpringFestivalDay],
	},
	{
		Start:  "2019/04/05",
		End:    "2019/04/07",
		ChName: chHolidays[TombSweepingDay],
		EnName: enHolidays[TombSweepingDay],
	},
	{
		Start:  "2019/04/29",
		End:    "2019/05/01",
		ChName: chHolidays[LaborDay],
		EnName: enHolidays[LaborDay],
	},
	{
		Start:  "2019/06/07",
		End:    "2019/06/09",
		ChName: chHolidays[DragonBoatFestivalDay],
		EnName: enHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2019/09/13",
		End:    "2019/09/15",
		ChName: chHolidays[MidAutumnFestivalDay],
		EnName: enHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2019/10/01",
		End:    "2019/10/07",
		ChName: chHolidays[NationalDay],
		EnName: enHolidays[NationalDay],
	},
}
var holiday2010 = []YearCollection{
	{
		Start:  "2018/12/30",
		End:    "2019/01/01",
		ChName: chHolidays[NewYearDay],
		EnName: enHolidays[NewYearDay],
	},
	{
		Start:  "2019/02/04",
		End:    "2019/02/10",
		ChName: chHolidays[SpringFestivalDay],
		EnName: enHolidays[SpringFestivalDay],
	},
	{
		Start:  "2019/04/05",
		End:    "2019/04/07",
		ChName: chHolidays[TombSweepingDay],
		EnName: enHolidays[TombSweepingDay],
	},
	{
		Start:  "2019/04/29",
		End:    "2019/05/01",
		ChName: chHolidays[LaborDay],
		EnName: enHolidays[LaborDay],
	},
	{
		Start:  "2019/06/07",
		End:    "2019/06/09",
		ChName: chHolidays[DragonBoatFestivalDay],
		EnName: enHolidays[DragonBoatFestivalDay],
	},
	{
		Start:  "2019/09/13",
		End:    "2019/09/15",
		ChName: chHolidays[MidAutumnFestivalDay],
		EnName: enHolidays[MidAutumnFestivalDay],
	},
	{
		Start:  "2019/10/01",
		End:    "2019/10/07",
		ChName: chHolidays[NationalDay],
		EnName: enHolidays[NationalDay],
	},
}
