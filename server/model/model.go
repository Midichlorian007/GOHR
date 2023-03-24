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
