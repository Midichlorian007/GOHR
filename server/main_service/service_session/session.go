package service_session

import (
	"GOHR/server/db"
	"GOHR/server/model"

	"github.com/gin-gonic/gin"
)

type SessionInterface interface {
	AddSession(ctx *gin.Context, session *model.Session)
	CheckSession(ctx *gin.Context, session string) int
}

type sessionStruct struct {
	db db.DbInterface
}

func New(db db.DbInterface) SessionInterface {
	return &sessionStruct{
		db: db,
	}
}

func (s *sessionStruct) AddSession(ctx *gin.Context, session *model.Session) {

	// all := u.db.GetAllUser(ctx)
	// if ctx.IsAborted() {
	// 	return nil
	// }

}

func (s *sessionStruct) CheckSession(ctx *gin.Context, session string) int {

	// all := u.db.GetAllUser(ctx)
	// if ctx.IsAborted() {
	// 	return nil
	// }

	return 0
}
