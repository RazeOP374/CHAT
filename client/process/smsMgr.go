package process

import (
	"GOproject/project1/chatroom/common/message"
	"encoding/json"
	"fmt"
)

func outputGruopMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("jsonunmarshal err", err)
		return
	}
	info := fmt.Sprintf("用户id:\t%d 说:\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
