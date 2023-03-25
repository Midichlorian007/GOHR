package model

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	//ETC
}

type HR struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Company  string `json:"company"`
	//ETC
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
