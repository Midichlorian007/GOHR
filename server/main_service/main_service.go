package main_service

import (
	"GOHR/server/db"
	"GOHR/server/main_service/service_course"
	"GOHR/server/main_service/service_hr"
	"GOHR/server/main_service/service_session"
	"GOHR/server/main_service/service_solve"
	"GOHR/server/main_service/service_users"
	"GOHR/server/model"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type MainServiceInterface interface {
	GetAllUsers(ctx *gin.Context) []*model.User
	GetProfile(ctx *gin.Context, session string) *model.Profile
	AddNewUser(ctx *gin.Context, user *model.SignUp)
	SignUser(ctx *gin.Context, user *model.SignIn)
}

type mainServiceStruct struct {
	Users   service_users.UsersInterface
	HR      service_hr.HRInterface
	Solve   service_solve.SolveInterface
	Course  service_course.CourseInterface
	Session service_session.SessionInterface
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

func (s *mainServiceStruct) GetProfile(ctx *gin.Context, session string) *model.Profile {
	// check session
	userID := s.Session.CheckSession(ctx, session)
	if ctx.IsAborted() {
		return nil
	}
	var profile model.Profile
	// Get user info
	profile.User = s.Users.GetUserByID(ctx, userID)
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

func (s *mainServiceStruct) AddNewUser(ctx *gin.Context, user *model.SignUp) {
	exists := s.Users.CheckUserExists(ctx, user.Name)
	if ctx.IsAborted() {
		return
	}
	if exists {
		errMsg := "error AddNewUser: user already exists " + user.Name
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New(errMsg)))
		return
	}
	// create new user
	s.Users.AddNewUser(ctx, user)
	if ctx.IsAborted() {
		return
	}
	// create session
	session := model.Session{}
	s.Session.AddSession(ctx, &session)
	if ctx.IsAborted() {
		return
	}
}

func (s *mainServiceStruct) SignUser(ctx *gin.Context, user *model.SignIn) {
	// get hashed password
	secret := s.Users.GetUserSecret(ctx, user.Name)
	if ctx.IsAborted() {
		return
	}
	// validate
	if err := bcrypt.CompareHashAndPassword([]byte(secret), []byte(user.Password)); err != nil &&
		errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) { // passwords did not match
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New("login or password wrong")))
		return
	}
	// create session
	session := model.Session{}
	s.Session.AddSession(ctx, &session)
	if ctx.IsAborted() {
		return
	}
}
