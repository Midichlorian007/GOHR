package service_solve

import (
	"GOHR/server/db"
	"GOHR/server/model"

	"github.com/gin-gonic/gin"
)

type SolveInterface interface {
	GetProblemSolves(ctx *gin.Context, id string) []*model.ProblemSolve
}

type solveStruct struct {
	db db.DbInterface
}

func New(db db.DbInterface) SolveInterface {
	return &solveStruct{
		db: db,
	}
}

func (s *solveStruct) GetProblemSolves(ctx *gin.Context, id string) []*model.ProblemSolve {

	// all := u.db.GetAllUser(ctx)
	// if ctx.IsAborted() {
	// 	return nil
	// }

	return nil
}
