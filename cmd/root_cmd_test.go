package cmd

import "testing"

func TestHandlerString(tests *testing.T) {

	var tt = []struct {
		startDate string
		endDate   string
	}{
		{
			startDate: "2006/01/02",
			endDate:   "2009/01/04",
		},
	}

	for _, t := range tt {
		handlerString(t.startDate)
		handlerString(t.endDate)
	}

}
