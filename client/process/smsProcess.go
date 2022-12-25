package process

import (
	"GOproject/project1/chatroom/client/utils"
	"GOproject/project1/chatroom/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("sendgroupMesERR", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("sendgroupMesERR")
		return
	}
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("send err")
		return
	}
	return
}
func (this *SmsProcess) SendToOne(content string, id int) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType
	var smsMes message.SmsMes
	smsMes.Id = id
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("sendgroupMesERR", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("sendgroupMesERR")
		return
	}
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("send err")
		return
	}
	return
}
func (this *SmsProcess) OutPut(mes *message.Message) (err error) {
	var SmsResMes message.SmsResMes
	err = json.Unmarshal([]byte(mes.Data), &SmsResMes)
	if err != nil {
		fmt.Println("u你Marshalerr", err)
		return
	}
	if SmsResMes.Id == 0 {
		outputGruopMes(mes)
	} else {
		outputoneMes(mes)
	}
	return
}
