package holiday

type FetcherParams struct {
	Return  string `form:"return" json:"return"`
	OffSet  int    `form:"offset" json:"off_set"`
	Limit   int    `form:"limit" json:"limit"`
	PerPage int    `form:"per_page" json:"per_page"`
	Page    int    `form:"page" json:"page"`
}

type FetcherByYearParams struct {
	Year      string `form:"year" json:"year"`
	WithCount string `form:"with_count" json:"with_count"`
}

type FetcherByMonthParams struct {
	Month     string `form:"month" json:"month"`
	WithCount string `form:"with_count" json:"with_count"`
}
