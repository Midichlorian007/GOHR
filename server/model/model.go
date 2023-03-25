package model

import "time"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type Session struct {
	ID     string    `json:"id"`
	UserID string    `json:"user_id"`
	Expire time.Time `json:"expire"`
}

type Company struct {
	BIN  string `json:"bin"`
	Name string `json:"name"`
}

type ProblemSolve struct {
	Name   string `json:"name"`
	Total  int    `json:"total"`
	Easy   int    `json:"easy"`
	Medium int    `json:"medium"`
	Hard   int    `json:"hard"`
}

type Course struct {
	ID string `json:"id"`
}

type Profile struct {
	User          *User           `json:"user"`
	ProblemSolves []*ProblemSolve `json:"problem_solves"`
	Courses       []*Course       `json:"courses"`
}
