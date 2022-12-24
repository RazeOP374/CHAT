package process

import (
	"GOproject/project1/chatroom/common/message"
	"GOproject/project1/chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGruopMes(mes *message.Message) {

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		return
	}
	for id, up := range userMgr.onlineUser {
		if id == smsMes.UserId {
			continue
		}
		this.SendMestoUser(data, up.Conn)

	}

}
func (this *SmsProcess) SendMestoUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发失败")
	}
}
