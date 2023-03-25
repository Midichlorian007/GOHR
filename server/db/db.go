package db

import (
	"GOHR/server/model"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DbInterface interface {
	GetAllUser(ctx *gin.Context) []*model.User
	AddUser(ctx *gin.Context, user *model.SignUp, secret string)
	GetUser(ctx *gin.Context, login string) *model.User
	GetUserSecret(ctx *gin.Context, login string) string
}

type dbStruct struct {
	db *sql.DB
}

func New(dbPath, dbDriver string) (DbInterface, func() error) {

	db, err := sql.Open(dbDriver, dbPath)
	if err != nil {
		createDB(dbPath) //To create SQLite database
		db, _ = sql.Open(dbDriver, dbPath)
		createTable(db) //To create table 'users'
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &dbStruct{db: db}, db.Close
}

func (d *dbStruct) GetAllUser(ctx *gin.Context) []*model.User {
	qry := `select * from users;`
	rows, err := d.db.Query(qry)
	if err != nil {
		errMsg := "error db GetAllUser query: " + err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New(errMsg)))
		return nil
	}
	defer rows.Close()

	allUsers := []*model.User{}

	for rows.Next() {
		nextUser := model.User{}
		err = rows.Scan(&nextUser.ID, &nextUser.Name, &nextUser.LastName, &nextUser.Email)
		if err != nil {
			errMsg := "error db GetAllUser scan: " + err.Error()
			ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New(errMsg)))
			return nil
		}
		allUsers = append(allUsers, &nextUser)
	}
	return allUsers
}

func (d *dbStruct) AddUser(ctx *gin.Context, user *model.SignUp, secret string) {

	qry := `INSERT INTO users(name, last_name, email ) VALUES (?, ?, ?); `
	statement, err := d.db.Prepare(qry)
	if err != nil {
		errMsg := "error db AddUser: " + err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New(errMsg)))
		return
	}

	_, err = statement.Exec(user.Name, user.LastName, user.Email)
	if err != nil {
		errMsg := "error db AddUser: " + err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New(errMsg)))
		return
	}
}

func (d *dbStruct) GetUser(ctx *gin.Context, login string) *model.User {
	var user model.User

	//TODO IMPLEMENT ME

	return &user
}

func (d *dbStruct) GetUserSecret(ctx *gin.Context, login string) string {

	//TODO IMPLEMENT ME

	return ""
}
