package service_hr

import (
	"GOHR/server/db"
	"GOHR/server/model"
	"github.com/gin-gonic/gin"
)

type HRInterface interface {
	AddNewHR(ctx *gin.Context, newHr model.HR)
}

type hrStruct struct {
	db db.DbInterface
}

func New(db db.DbInterface) HRInterface {
	return &hrStruct{
		db: db,
	}
}

func (hr *hrStruct) AddNewHR(ctx *gin.Context, newHr model.HR) {
	hr.db.AddHR(ctx, newHr)
}
