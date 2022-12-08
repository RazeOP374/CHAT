package main

import (
	"fmt"
	"net"
)

// 读取客户端发送信息,并验证长度
/*func readpkg(conn net.Conn) (mes common.Message, err error) {
	buf := make([]byte, 4096)
	fmt.Println("读取中")
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

// 服务器端data长度验证，防止丢包
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
}*/

// 服务器登陆验证
/*func severProcessLogin(conn net.Conn, mes *common.Message) (err error) {
	var loginmes common.Loginmes
	err = json.Unmarshal([]byte(mes.Data), &loginmes)
	if err != nil {
		panic(err)
		return
	}
	var resMes common.Message
	resMes.Type = common.LoginResmestype
	var loginResMes common.LoginResmes
	if loginmes.Userid == 123456 && loginmes.Userpwd == "abco123" {
		loginResMes.Code = 200
		loginResMes.Error = "登录成功"
	} else {
		loginResMes.Code = 418
		loginResMes.Error = "请重新注册"
	}
	data, err := json.Marshal(loginResMes)
	if err != nil {
		panic(err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		panic(err)
		return
	}
	err = writepkg(conn, data)
	return
}
*/
// 服务器端功能处理
/*func serverProcessMes(conn net.Conn, mes *common.Message) (err error) {
	switch mes.Type {
	case common.Loginmestype:
		err = severProcessLogin(conn, mes)
	case common.RegisterMestype:
	default:
		panic(err)
	}
	return
}*/

// 协程接口处理客户端消息功能
func mainprocessa(conn net.Conn) {
	defer conn.Close()
	processor := &Processor{
		Conn: conn,
	}
	err := processor.lo()
	if err != nil {
		//	panic(err)
		return
	}
}

// 服务器端连接监听，接收客户端消息
func main() {
	fmt.Println("服务器8889")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		panic(err)
		return
	}
	for {
		fmt.Println("等待中")
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go mainprocessa(conn)
	}
}
