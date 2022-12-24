package main

import (
	"GOproject/project1/chatroom/common/message"
	process2 "GOproject/project1/chatroom/server/process"
	"GOproject/project1/chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {

	fmt.Println(mes)
	switch mes.Type {
	case message.Loginmestype:
		//处理登录
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.SeverProcessLogin(mes)
	case message.RegisterMestype:
		//处理注册
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.SeverProcessRegister(mes)
	case message.SmsMesType:
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGruopMes(mes)
	default:
		fmt.Println("无法处理消息类型")
	}
	return
}
func (this *Processor) lo() (err error) {

	//循环客户端消息
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出")
				return err
			} else {
				return err
			}
			//panic(err)
		}
		err = this.ServerProcessMes(&mes)
		if err != nil {
			//panic(err)
			return err
		}
	}
}
