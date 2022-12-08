package utils

import (
	"GOproject/project1/chatroom/common"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [4096]byte
}

func (this *Transfer) Readpkg( /*conn net.Conn*/ ) (mes common.Message, err error) {
	//buf := make([]byte, 4096)
	fmt.Println("读取中")
	_, err = this.Conn.Read(this.Buf[:8])
	if err != nil {
		//	panic(err)
		return
	}
	var pkglen uint64
	pkglen = binary.BigEndian.Uint64(this.Buf[0:8])
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
func (this *Transfer) Writepkg(data []byte) (err error) {
	var pkglen uint64
	pkglen = uint64(len(data))
	//	var bytes [8]byte
	binary.BigEndian.PutUint64(this.Buf[0:8], pkglen)
	n, err := this.Conn.Write(this.Buf[:8])
	if n != 8 || err != nil {
		panic(err)
		return
	}
	n, err = this.Conn.Write(data)
	if n != int(pkglen) || err != nil {
		panic(err)
		return
	}
	return
}
