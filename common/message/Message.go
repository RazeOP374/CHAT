package message

const (
	Loginmestype       = "Loginmes"
	LoginResmestype    = "LoginResmes"
	RegisterMestype    = "RegisterMes"
	RegisterResMestype = "RegisterResMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
type LoginResmes struct {
	Code    int `json:"code"`
	UserIds []int
	Error   string `json:"error"`
}
type RegisterMes struct {
	User User `json:"user"`

	//registerpwd string `json:"registerpwd"`
}
type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
