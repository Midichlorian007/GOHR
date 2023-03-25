package main_service

import (
	"GOHR/server/db"
	"GOHR/server/main_service/service_course"
	"GOHR/server/main_service/service_hr"
	"GOHR/server/main_service/service_solve"
	"GOHR/server/main_service/service_users"
	"GOHR/server/model"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainServiceInterface interface {
	GetAllUsers(ctx *gin.Context) []*model.User
	GetProfile(ctx *gin.Context) *model.Profile
	AddNewUser(ctx *gin.Context, user model.User)
}

type mainServiceStruct struct {
	Users  service_users.UsersInterface
	HR     service_hr.HRInterface
	Solve  service_solve.SolveInterface
	Course service_course.CourseInterface
	// TODO Тут добовляем необходимые сервисы для выполнения бизнес логики, но не БД
}

func New(db db.DbInterface) MainServiceInterface {
	return &mainServiceStruct{
		// TODO Тут имплементация необходимых сервисов
		Users: service_users.New(db),
		HR:    service_hr.New(db),
		Solve: service_solve.New(db),
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

func (s *mainServiceStruct) GetProfile(ctx *gin.Context) *model.Profile {
	var profile model.Profile
	// Get user info
	profile.User = s.Users.GetUser(ctx)
	if ctx.IsAborted() {
		return nil
	}
	// get problem solves
	profile.ProblemSolves = s.Solve.GetProblemSolves(ctx, profile.User.ID)
	if ctx.IsAborted() {
		return nil
	}
	// get current courses
	profile.Courses = s.Course.GetCourses(ctx, profile.User.ID)
	if ctx.IsAborted() {
		return nil
	}
	return &profile
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
