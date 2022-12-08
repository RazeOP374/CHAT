package main

import (
	"GOproject/project1/chatroom/common"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func login(id int, password string) (err error) {
	fmt.Printf("账户=%v,密码=%s", id, password)
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		panic(err)
		return
	}
	defer conn.Close()
	var mes common.Message
	mes.Type = common.Loginmestype
	var loginMes common.Loginmes
	loginMes.Userid = id
	loginMes.Userpwd = password
	data, err := json.Marshal(loginMes)
	if err != nil {
		panic(err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		panic(err)
		return
	}
	var pkglen uint64
	pkglen = uint64(len(data))
	var bytes [8]byte
	binary.BigEndian.PutUint64(bytes[0:8], pkglen)
	n, err := conn.Write(bytes[:8])
	if n != 8 || err != nil {
		panic(err)
		return
	}
	_, err = conn.Write(data)
	if err != nil {
		panic(err)
	}
	mes, err = readpkg(conn)
	if err != nil {
		return
	}
	var loginResMes common.LoginResmes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功")
	} else if loginResMes.Code == 418 {
		fmt.Println(loginResMes.Error)
	}
	return
}
