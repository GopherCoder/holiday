package holiday

import "github.com/gin-gonic/gin"

func RegisterHoliday(r *gin.RouterGroup) {

	var controller controllerHoliday

	r.POST("/creator", controller.PostHolidayHandler)
	r.PATCH("/patcher", controller.UpdateHolidayHandler)
	r.GET("/fetcher", controller.GetHolidayHandler)
	r.DELETE("/deleter", controller.DeleteHolidayHandler)
	r.GET("/fetcher/by/years", controller.GetHolidayByYearHandler)
	r.GET("/fetcher/by/months", controller.GetHolidayByMonthHandler)
	r.GET("/fetcher/by/days", controller.GetHolidayByDaysHandler)

}
