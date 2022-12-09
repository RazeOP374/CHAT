package common

const (
	Loginmestype    = "Loginmes"
	LoginResmestype = "LoginResmes"
	RegisterMestype = "RegisterMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
type Loginmes struct {
	Userid   int    `json:"userid"`
	Userpwd  string `json:"userpwd"`
	UserName string `json:"userName"`
}
type LoginResmes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
type RegisterMes struct {
	User User `json:"user"`

}
type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
