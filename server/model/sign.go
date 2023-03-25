package model

type SignIn struct {
	Name     string `form:"inputName" json:"name"`
	Password string `form:"inputPass" json:"password"`
}

func (s SignIn) Valid() bool {
	if s.Name == "" {
		return false
	}
	if s.Password == "" {
		return false
	}
	return true
}

type SignUp struct {
	Name     string `form:"inputNameup" json:"name"`
	LastName string `form:"inputLastName" json:"last_name"`
	Password string `form:"inputPassup" json:"password"`
	Email    string `form:"inputMail" json:"email"`
}