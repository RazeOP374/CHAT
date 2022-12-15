package process

import (
	"GOproject/project1/chatroom/client/utils"
	"fmt"
	"net"
	"os"
)

func ShowMenu() {
	for {
		fmt.Println("----------登陆成功----------")
		fmt.Println("----------1.显示用户在线列表----------")
		fmt.Println("----------2.发送消息----------")
		fmt.Println("----------3.历史消息----------")
		fmt.Println("----------4.退出系统----------")
		fmt.Println("请选择")
		var key int
		fmt.Scan(&key)
		switch key {
		case 1:
			fmt.Println()
		case 2:
			fmt.Println()
		case 3:
			fmt.Println()
		case 4:
			fmt.Println("退出")
			os.Exit(4)
		default:
			fmt.Println("输出有误")
		}
	}
}
func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端等待读取中")
		mes, err := tf.ReadPkg()
		if err != nil {
			return
		}
		fmt.Println(mes)
	}
}
