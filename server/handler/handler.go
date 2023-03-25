package handler

import (
	"GOHR/server/main_service"
	"GOHR/server/model"
	"errors"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HandlerInterface interface {
	GetAllUsers(ctx *gin.Context)
	AddNewHR(ctx *gin.Context)
	Index(ctx *gin.Context)
	SignIn(ctx *gin.Context)
	SignUp(ctx *gin.Context)
	Logout(ctx *gin.Context)
	// Profile(ctx *gin.Context)
}

func New(main_service main_service.MainServiceInterface) HandlerInterface {
	return &handlerStuct{
		service: main_service,
	}
}

type handlerStuct struct {
	service main_service.MainServiceInterface
}

func (h *handlerStuct) GetAllUsers(ctx *gin.Context) {
	res := h.service.GetAllUsers(ctx)
	if ctx.IsAborted() {
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *handlerStuct) SignUp(ctx *gin.Context) {
	user := model.SignUp{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		errMsg := "error AddNewUser bind body: " + err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New(errMsg)))
		return
	}

	h.service.AddNewUser(ctx, &user)
	if ctx.IsAborted() {
		return
	}
	ctx.Status(http.StatusCreated)
}

func (h *handlerStuct) AddNewHR(ctx *gin.Context) {
	// TODO IMPLEMENT ME
	ctx.Status(http.StatusCreated)
}

func (h *handlerStuct) Index(ctx *gin.Context) {
	session, _ := ctx.Get("session")
	profile := h.service.GetProfile(ctx, session.(string))
	if ctx.IsAborted() {
		ctx.HTML(http.StatusUnauthorized, "login.html", gin.H{})
		return
	}
	ctx.HTML(http.StatusOK, "profile.html", gin.H{"profile": profile})
}

func (h *handlerStuct) Logout(ctx *gin.Context) {

	user, _ := ctx.Cookie(viper.GetString("session_user"))

	ctx.SetCookie(user, "", -1, "http://127.0.0.1:9090/", "/", false, true)

	opt := gin.H{
		"user": "GUEST",
	}
	ctx.HTML(http.StatusOK, "login.html", opt)
}

func (h *handlerStuct) SignIn(ctx *gin.Context) {
	user := model.SignIn{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		errMsg := "error SignIn bind body: " + err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New(errMsg)))
		return
	}
	// check for not null values and validate values
	if !user.Valid() {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New("empty fields or not valid")))
		return
	}
	// sign user
	h.service.SignUser(ctx, &user)
	if ctx.IsAborted() {
		return
	}
	ctx.Status(http.StatusCreated)
}
