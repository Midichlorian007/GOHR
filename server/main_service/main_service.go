package main_service

import (
	"GOHR/server/db"
	"GOHR/server/main_service/service_hr"
	"GOHR/server/main_service/service_users"
	"GOHR/server/model"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MainServiceInterface interface {
	GetAllUsers(ctx *gin.Context) []*model.User
	AddNewUser(ctx *gin.Context, user model.User)
}

type mainServiceStruct struct {
	Users service_users.UsersInterface
	HR    service_hr.HRInterface
	// TODO Тут добовляем необходимые сервисы для выполнения бизнес логики, но не БД
}

func New(db db.DbInterface) MainServiceInterface {
	return &mainServiceStruct{
		// TODO Тут имплементация необходимых сервисов
		Users: service_users.New(db),
		HR:    service_hr.New(db),
	}
}

// TODO Вот тут пишем бизнес логику сервиса и вызываем то что нам надо из других сервисов
// TODO Они называется так же как и хендлеры
func (s *mainServiceStruct) GetAllUsers(ctx *gin.Context) []*model.User {
	allUsers := s.Users.GetAllUsers(ctx)
	if ctx.IsAborted() {
		return nil
	}

	return allUsers
}

func (s *mainServiceStruct) AddNewUser(ctx *gin.Context, user model.User) {

	exists := s.Users.CheckUserExists(ctx, user)
	if ctx.IsAborted() {
		return
	}

	if exists {
		errMsg := "error AddNewUser: user already exists " + user.Name
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New(errMsg)))
		return
	}

	s.Users.AddNewUser(ctx, user)
	if ctx.IsAborted() {
		return
	}

}
