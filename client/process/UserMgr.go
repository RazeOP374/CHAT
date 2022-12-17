package process

import (
	"GOproject/project1/chatroom/common/message"
	"fmt"
)

var onlineUsers = make(map[int]*message.User, 10)

func outputonlineUser() {
	for id, _ := range onlineUsers {
		fmt.Println("用户id%t", id)
	}
}
func updateUserStatus(notifyUserstatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserstatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserstatusMes.UserId,
		}

	}

	user.UserStatus = notifyUserstatusMes.Status
	onlineUsers[notifyUserstatusMes.UserId] = user

	outputonlineUser()
}
