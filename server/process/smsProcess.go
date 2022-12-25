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
func (this *SmsProcess) SendMes(mes *message.Message) (err error) {
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("unmarshal err", err)
		return
	}
	if smsMes.Id == 0 {
		this.SendGruopMes(mes)
	} else {
		this.SendMesToOne(&smsMes)
	}
	return
}
func (this *SmsProcess) SendMesToOne(smsRes *message.SmsMes) (err error) {
	up := userMgr.onlineUser[smsRes.Id]
	var smsResMes message.SmsResMes
	smsResMes.Content = smsRes.Content
	smsResMes.User = smsRes.User
	smsResMes.Id = smsRes.Id
	data, err := json.Marshal(smsResMes)
	if err != nil {
		fmt.Println("smsresmes marshalerr", err)
		return
	}
	mes := message.Message{
		Type: message.SmsResMesType,
		Data: string(data),
	}
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("marshal err", err)
		return
	}
	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("write err", err)
		return
	}
	return
}
