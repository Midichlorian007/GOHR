package service_users

import (
	"GOHR/server/db"
	"GOHR/server/model"
	"github.com/gin-gonic/gin"
)

type UsersInterface interface {
	AddNewUser(ctx *gin.Context, user model.User)
	CheckUserExists(ctx *gin.Context, user model.User) bool
	GetAllUsers(ctx *gin.Context) []*model.User
}

type usersStruct struct {
	db db.DbInterface
}

func New(db db.DbInterface) UsersInterface {
	return &usersStruct{
		db: db,
	}
}

func (u *usersStruct) CheckUserExists(ctx *gin.Context, user model.User) bool {
	ok := u.db.CheckUserExists(ctx, user)

	return ok
}
func (u *usersStruct) AddNewUser(ctx *gin.Context, user model.User) {
	u.db.AddUser(ctx, user)
}

func (u *usersStruct) GetAllUsers(ctx *gin.Context) []*model.User {

	all := u.db.GetAllUser(ctx)
	if ctx.IsAborted() {
		return nil
	}

	return all
}
