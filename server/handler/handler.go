package handler

import (
	"GOHR/server/main_service"
	"GOHR/server/model"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerInterface interface {
	GetAllUsers(ctx *gin.Context)
	AddNewUser(ctx *gin.Context)
	AddNewHR(ctx *gin.Context)
	Index(ctx *gin.Context)
	Login(ctx *gin.Context)
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

func (h *handlerStuct) AddNewUser(ctx *gin.Context) {
	user := model.User{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		errMsg := "error AddNewUser bind body: " + err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New(errMsg)))
		return
	}

	h.service.AddNewUser(ctx, user)
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

	if "user" == "" {
		opt := gin.H{
			"text": "GOHR text",
		}

		ctx.HTML(http.StatusOK, "login.html", opt)
	}

	opt := gin.H{
		"text": "GOHR text",
	}
	ctx.HTML(http.StatusOK, "profile.html", opt)
}

func (h *handlerStuct) Login(ctx *gin.Context) {

	opt := gin.H{
		"text": "GOHR text",
	}
	ctx.HTML(http.StatusOK, "profile.html", opt)
}
