package process

import (
	"GOproject/project1/chatroom/client/utils"
	"GOproject/project1/chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ShowMenu() {
	for {
		fmt.Println("----------登陆成功----------")
		fmt.Println("----------1.广播消息----------")
		fmt.Println("----------2.发送消息----------")
		fmt.Println("----------3.历史消息----------")
		fmt.Println("----------4.退出系统----------")
		fmt.Println("请选择")
		var key int
		var content string
		smsprocess := &SmsProcess{}
		fmt.Scan(&key)
		switch key {
		case 1:
			fmt.Println("广播消息")
			fmt.Scan(&content)
			smsprocess.SendGroupMes(content)
		case 2:
			outputonlineUser()
			fmt.Println("私聊")
			var id int
			fmt.Println("请输入ID")
			fmt.Scan(&id)
			fmt.Println("-----私聊-----")
			for {
				fmt.Println("请输入内容")
				fmt.Scan(&content)
				smsprocess.SendToOne(content, id)
			}
		case 3:

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
		switch mes.Type {
		case message.NotifyUserStatusMestype: // 有人上线了

			//1. 取出.NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//2. 把这个用户的信息，状态保存到客户map[int]User中
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType: //有人群发消息
			outputGruopMes(&mes)
		case message.SmsResMesType:
			SmsProcess := SmsProcess{}
			SmsProcess.OutPut(&mes)
		default:
			fmt.Println("返回无法识别")
		}
	}
}
