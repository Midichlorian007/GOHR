package service_course

import (
	"GOHR/server/db"
	"GOHR/server/model"

	"github.com/gin-gonic/gin"
)

type CourseInterface interface {
	GetCourses(ctx *gin.Context, id string) []*model.Course
}

type courseStruct struct {
	db db.DbInterface
}

func New(db db.DbInterface) CourseInterface {
	return &courseStruct{
		db: db,
	}
}

func (c *courseStruct) GetCourses(ctx *gin.Context, id string) []*model.Course {

	// all := u.db.GetAllUser(ctx)
	// if ctx.IsAborted() {
	// 	return nil
	// }

	return nil
}
