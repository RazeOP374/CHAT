package process

import (
	"GOproject/project1/chatroom/common/message"
	"GOproject/project1/chatroom/server/model"
	"GOproject/project1/chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) SeverProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail 1")
		return
	}
	//先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMestype
	var registerResMes message.RegisterResMes
	//1.使用model.MyUserDao 到redis去验证
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "未知错误"
		}
	} else {
		registerResMes.Code = 200
	}
	data, err := json.Marshal(registerResMes)
	if err != nil {
		//panic(err)
		fmt.Println("json.Marshal fail 2")
		return
	}
	resMes.Data = string(data)
	//对resMes进行序列化，发送
	data, err = json.Marshal(resMes)
	if err != nil {
		//panic(err)
		fmt.Println("json.Marshal fail 3")
		return
	}
	tf := utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
func (this *UserProcess) SeverProcessLogin(mes *message.Message) (err error) {
	var loginmes message.LoginMes
	// 先从mes 中取出 mes.Data ，并直接反序列化成LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginmes)
	if err != nil {
		//panic(err)
		fmt.Println("json.Unmarshal fail err 1")
		return
	}
	var resMes message.Message
	resMes.Type = message.LoginResmestype
	var loginResMes message.LoginResmes
	user, err := model.MyUserDao.Login(loginmes.UserId, loginmes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 404
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 400
			loginResMes.Error = err.Error()
		}
	} else {
		loginResMes.Code = 200
		fmt.Println(user, "成功")
	}
	/*if loginmes.Userid == 123456 && loginmes.Userpwd == "123" {
		loginResMes.Code = 200
		loginResMes.Error = "登录成功"
	} else {
		loginResMes.Code = 418
		loginResMes.Error = "请重新注册"
	}*/
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail 2")
		//panic(err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail 3")

		//panic(err)
		return
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
