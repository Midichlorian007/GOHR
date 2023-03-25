package middleware

import (
	"GOHR/server/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func CheckSession(ctx *gin.Context) {
	session, err := ctx.Cookie(viper.GetString("cookies.session_user"))
	if err != nil && session == "" {
		ctx.HTML(http.StatusUnauthorized, "login.html", gin.H{})
		ctx.Abort()
	}
	ctx.Set(model.SessionKey, session)
}
