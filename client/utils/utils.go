package utils

import (
	"GOproject/project1/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 4096)
	fmt.Println("读取客户端发送的数据...")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//panic(err)
		return
	}
	var pkglen uint32
	pkglen = binary.BigEndian.Uint32(this.Buf[0:4])
	n, err := this.Conn.Read(this.Buf[:pkglen])
	if n != int(pkglen) || err != nil {
		//panic(err)
		return
	}
	err = json.Unmarshal(this.Buf[:pkglen], &mes)
	if err != nil {
		return
	}
	return
}

// 服务器端data长度验证，防止丢包
func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkglen uint32
	pkglen = uint32(len(data))
	//	var bytes [8]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkglen)
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("connwirtefail")
		//panic(err)
		return
	}
	n, err = this.Conn.Write(data)
	if n != int(pkglen) || err != nil {
		fmt.Println("conwritefail2")
		//panic(err)
		return
	}
	return
}
