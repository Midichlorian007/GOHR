package service_users

import (
	"GOHR/server/db"
	"GOHR/server/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UsersInterface interface {
	AddNewUser(ctx *gin.Context, user *model.SignUp)
	CheckUserExists(ctx *gin.Context, login string) bool
	GetAllUsers(ctx *gin.Context) []*model.User
	GetUser(ctx *gin.Context, login string) *model.User
	GetUserSecret(ctx *gin.Context, login string) string
}

type usersStruct struct {
	db db.DbInterface
}

func New(db db.DbInterface) UsersInterface {
	return &usersStruct{
		db: db,
	}
}

func (u *usersStruct) GetUser(ctx *gin.Context, login string) *model.User {
	user := u.db.GetUser(ctx, login)
	if ctx.IsAborted() {
		return nil
	}
	return user
}

func (u *usersStruct) CheckUserExists(ctx *gin.Context, login string) bool {
	user := u.db.GetUser(ctx, login)
	if ctx.IsAborted() {
		return false
	}
	return user != nil
}

func (u *usersStruct) AddNewUser(ctx *gin.Context, user *model.SignUp) {
	// crypt password
	secret, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(err))
		return
	}
	// add user
	u.db.AddUser(ctx, user, string(secret))
}

func (u *usersStruct) GetAllUsers(ctx *gin.Context) []*model.User {

	all := u.db.GetAllUser(ctx)
	if ctx.IsAborted() {
		return nil
	}

	return all
}

func (u *usersStruct) GetUserSecret(ctx *gin.Context, login string) string {
	return u.db.GetUserSecret(ctx, login)
}