package main

import (
	"GOproject/project1/chatroom/common"
	"GOproject/project1/chatroom/sever/process"
	"GOproject/project1/chatroom/sever/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) ServerProcessMes(mes *common.Message) (err error) {
	switch mes.Type {
	case common.Loginmestype:
		//处理登录
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.SeverProcessLogin(mes)
	case common.RegisterMestype:
		//处理注册
	default:
		panic(err)
	}
	return
}
func (this *Processor) lo() (err error) {

	//循环客户端消息
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.Readpkg()
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
			panic(err)
			return err
		}
	}
}
