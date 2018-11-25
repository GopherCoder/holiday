package holiday

import (
	"fmt"
	"holiday/handler"
	"holiday/model"
	"holiday/pkg/controller"
	"holiday/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controllerHoliday struct {
}

func (c controllerHoliday) PostHolidayHandler(context *gin.Context) {

}
func (c controllerHoliday) GetHolidayHandler(context *gin.Context) {}
func (c controllerHoliday) GetHolidayByYearHandler(context *gin.Context) {

	var params FetcherByYearParams

	if err := controller.CheckQuery(context, &params); err != nil {
		return
	}

	fmt.Println(params)

	var info = logger.Info{
		Package: "holiday",
		Action:  "get tuples",
		Message: "get tuples fail",
	}

	var holidays []model.Holiday
	if dbError := model.PostgresDB.DB.Where("end_year = ?", params.Year).
		Find(&holidays).Error; dbError != nil {
		logger.DebugCtx(context, info)
		return
	}

	if params.WithCount == "true" {
		var Data []model.BasicSerialization
		for _, one := range holidays {
			Data = append(Data, one.Basic())
		}
		var response = handler.Response{
			Code:    http.StatusOK,
			Message: "get holiday by year success",
			Data:    Data,
		}
		handler.SendResponse(context, response)
	} else {
		var response = handler.Response{
			Code:    http.StatusOK,
			Message: "get holiday by year success",
			Data:    holidays,
		}
		handler.SendResponse(context, response)

	}
}
func (c controllerHoliday) GetHolidayByMonthHandler(context *gin.Context) {

	var params FetcherByMonthParams
	if err := controller.CheckQuery(context, &params); err != nil {
		return
	}

	var info = logger.Info{
		Package: "holiday",
		Action:  "get holiday record",
		Message: "get holiday record fail",
	}

	var holidays []model.Holiday
	if params.Month != "" {
		if dbError := model.PostgresDB.DB.Where("end_month = ?", params.Month).
			Find(&holidays).Error; dbError != nil {
			logger.InfoCtx(context, info)
			return
		}
	}

	if params.WithCount != "true" {
		var response = handler.Response{
			Code:    http.StatusOK,
			Message: "get holiday by month success",
			Data:    holidays,
		}
		handler.SendResponse(context, response)
	} else {
		var datas []model.BasicSerialization
		for _, one := range holidays {
			datas = append(datas, one.Basic())
		}
		var response = handler.Response{
			Code:    http.StatusOK,
			Message: "get holiday by month success",
			Data:    datas,
		}
		handler.SendResponse(context, response)
	}

}
func (c controllerHoliday) GetHolidayByDaysHandler(context *gin.Context) {}
func (c controllerHoliday) UpdateHolidayHandler(context *gin.Context)    {}
func (c controllerHoliday) DeleteHolidayHandler(context *gin.Context)    {}
