package process

import (
	"GOproject/project1/chatroom/client/utils"
	"GOproject/project1/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
}

func (this *UserProcess) Register(userId int,
	userPwd string, userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("netdial err")
		//		panic(err)
		return
	}
	defer conn.Close()
	var mes message.Message
	mes.Type = message.RegisterMestype
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("jsonmarshal err")
		//panic(err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("jsonmarshal err")
		//	panic(err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册失效")
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("Read err")
		return
	}
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	fmt.Printf("账户=%v,密码=%s", userId, userPwd)
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("netdailerr")
		//		panic(err)
		return
	}
	defer conn.Close()
	var mes message.Message
	mes.Type = message.Loginmestype
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("jsonmarshal err")
		//panic(err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("jsonmarshal err")
		//panic(err)
		return
	}
	var pkglen uint32
	pkglen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkglen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("connwrite err")
		//panic(err)
		return
	}
	fmt.Println("客户端发送消息")
	_, err = conn.Write(data)
	if err != nil {
		//	panic(err)
		return
	}
	tf := utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		return
	}
	var loginResMes message.LoginResmes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		for _, id := range loginResMes.UserIds {
			if id == userId {
				continue
			}
			fmt.Printf("用户id:%d\t", id)
		}
		fmt.Print("\n\n")
		go serverProcessMes(conn)
		ShowMenu()
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}
