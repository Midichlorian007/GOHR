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
	AddSession(ctx *gin.Context, session *model.Session)
	GetSession(ctx *gin.Context, id string) *model.Session
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

	qry := `SELECT id, name, last_name, email, role
	FROM users AS u
	WHERE u.name = ?;`

	statement, err := d.db.Prepare(qry)
	if err != nil {
		errMsg := "error db GetUser: " + err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New(errMsg)))
		return nil
	}

	err = statement.QueryRow(login).Scan(
		&user.ID,
		&user.Name,
		&user.LastName,
		&user.Email,
		&user.Role,
	)

	if err != nil {
		errMsg := "error db GetUser: " + err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New(errMsg)))
		return nil
	}

	return &user
}

func (d *dbStruct) GetUserSecret(ctx *gin.Context, login string) string {
	var secret string
	qry := `SELECT secret
	FROM users AS u
	WHERE u.name = ?;`

	statement, err := d.db.Prepare(qry)
	if err != nil {
		errMsg := "error db GetUserSecret: " + err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New(errMsg)))
		return ""
	}

	if err = statement.QueryRow(login).Scan(&secret); err != nil {
		errMsg := "error db GetUserSecret: " + err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New(errMsg)))
		return ""
	}

	return secret
}

func (d *dbStruct) AddSession(ctx *gin.Context, session *model.Session) {

	qry := `INSERT INTO sessions (id, user, expire)
	VALUES (?, ?, ?);`

	statement, err := d.db.Prepare(qry)
	if err != nil {
		errMsg := "error db AddSession: " + err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New(errMsg)))
		return
	}

	_, err = statement.Exec(session.ID, session.User, session.Expire)
	if err != nil {
		errMsg := "error db AddSession: " + err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New(errMsg)))
		return
	}
}

func (d *dbStruct) GetSession(ctx *gin.Context, id string) *model.Session {
	var session model.Session
	qry := `SELECT id, user, expire
	FROM sessions AS s
	WHERE s.id = ?;`

	statement, err := d.db.Prepare(qry)
	if err != nil {
		errMsg := "error db GetUserSecret: " + err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New(errMsg)))
		return nil
	}

	if err = statement.QueryRow(id).Scan(&session.ID, &session.User, &session.Expire); err != nil {
		errMsg := "error db GetUserSecret: " + err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Error(errors.New(errMsg)))
		return nil
	}

	return &session
}
