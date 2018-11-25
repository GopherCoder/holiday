package controller

import (
	"holiday/pkg/errors"

	"github.com/gin-gonic/gin"
)

func CheckQuery(context *gin.Context, value interface{}) error {
	if err := context.ShouldBindQuery(value); err != nil {
		return errors.ErrCheckQuery
	}
	return nil
}

func CheckForm(context *gin.Context, value interface{}) error {
	if err := context.ShouldBind(value); err != nil {
		return errors.ErrCheckForm
	}
	return nil
}

func CheckBody(context *gin.Context, value interface{}) error {
	if err := context.ShouldBindJSON(value); err != nil {
		return errors.ErrCheckBody
	}
	return nil
}
