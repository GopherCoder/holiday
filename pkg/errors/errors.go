package errors

import "fmt"

type HolidayError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Action  string `json:"action"`
}

func (h *HolidayError) Error() string {
	return fmt.Sprintf("Code=%d, Action=%s, Message=%s", h.Code, h.Action, h.Message)
}
