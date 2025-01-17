package dto

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type ReturnToken struct {
	Token *Token
	Error error
}

type ReturnConfirmation struct {
	Confirmation bool
	Error        error
}
