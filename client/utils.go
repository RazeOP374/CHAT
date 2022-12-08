package main

import (
	"GOproject/project1/chatroom/common"
	"encoding/binary"
	"encoding/json"
	"net"
)

func readpkg(conn net.Conn) (mes common.Message, err error) {
	buf := make([]byte, 4096)
	_, err = conn.Read(buf[:8])
	if err != nil {
		//	panic(err)
		return
	}
	var pkglen uint64
	pkglen = binary.BigEndian.Uint64(buf[0:8])
	n, err := conn.Read(buf[:pkglen])
	if n != int(pkglen) || err != nil {
		//panic(err)
		return
	}
	err = json.Unmarshal(buf[:pkglen], &mes)
	if err != nil {
		return
	}
	return
}
func writepkg(conn net.Conn, data []byte) (err error) {
	var pkglen uint64
	pkglen = uint64(len(data))
	var bytes [8]byte
	binary.BigEndian.PutUint64(bytes[0:8], pkglen)
	n, err := conn.Write(bytes[:8])
	if n != 8 || err != nil {
		panic(err)
		return
	}
	n, err = conn.Write(data)
	if n != int(pkglen) || err != nil {
		panic(err)
		return
	}
	return
}
