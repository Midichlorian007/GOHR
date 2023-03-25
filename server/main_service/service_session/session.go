package service_session

import (
	"GOHR/server/db"
	"GOHR/server/model"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SessionInterface interface {
	AddSession(ctx *gin.Context, session *model.Session)
	GetSession(ctx *gin.Context, session string) *model.Session
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
	s.db.AddSession(ctx, session)
}

func (s *sessionStruct) GetSession(ctx *gin.Context, id string) *model.Session {
	session := s.db.GetSession(ctx, id)
	if ctx.IsAborted() {
		return nil
	}
	if session.Expire.Before(time.Now()) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, ctx.Error(errors.New("session expired")))
		return nil
	}

	return session
}
