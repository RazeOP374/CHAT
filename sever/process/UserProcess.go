package process

import (
	"GOproject/project1/chatroom/common"
	"GOproject/project1/chatroom/sever/utils"
	"encoding/json"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) SeverProcessLogin(mes *common.Message) (err error) {
	var loginmes common.Loginmes
	err = json.Unmarshal([]byte(mes.Data), &loginmes)
	if err != nil {
		panic(err)
		return
	}
	var resMes common.Message
	resMes.Type = common.LoginResmestype
	var loginResMes common.LoginResmes
	if loginmes.Userid == 123456 && loginmes.Userpwd == "abco123" {
		loginResMes.Code = 200
		loginResMes.Error = "登录成功"
	} else {
		loginResMes.Code = 418
		loginResMes.Error = "请重新注册"
	}
	data, err := json.Marshal(loginResMes)
	if err != nil {
		panic(err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		panic(err)
		return
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.Writepkg(data)
	return
}
