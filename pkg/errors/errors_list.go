package errors

var (
	ErrCreate = &HolidayError{Code: 400, Action: "Create", Message: "Create tuple fail"}
	ErrQuery  = &HolidayError{Code: 400, Action: "Query", Message: "Query tuple fail"}
	ErrUpdate = &HolidayError{Code: 400, Action: "Update", Message: "Update tuple fail"}
	ErrDelete = &HolidayError{Code: 400, Action: "Delete", Message: "Delete tuple fail"}

	ErrCheckQuery = &HolidayError{Code: 400, Action: "Check", Message: "Check query fail"}
	ErrCheckForm  = &HolidayError{Code: 400, Action: "Check", Message: "Check form fail"}
	ErrCheckBody  = &HolidayError{Code: 400, Action: "Check", Message: "Check body fail"}
)
