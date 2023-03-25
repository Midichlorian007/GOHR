package service_hr

import (
	"GOHR/server/db"
)

type HRInterface interface {
}

type hrStruct struct {
	db db.DbInterface
}

func New(db db.DbInterface) HRInterface {
	return &hrStruct{
		db: db,
	}
}
