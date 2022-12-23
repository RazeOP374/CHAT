package model

import (
	"GOproject/project1/chatroom/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
