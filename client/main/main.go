package main

import (
	"GOproject/project1/chatroom/client/process"
	"fmt"
	"os"
)

var userId int
var userPwd string
var userName string

func main() {
	var key int
	//var loop = true
	for true {
		fmt.Println("-------欢迎登录聊天系统-------")
		fmt.Println("--------1.登录聊天室---------")
		fmt.Println("---------2.注册用户---------")
		fmt.Println("--------3.退出聊天室---------")
		fmt.Println("请输入1-3")
		fmt.Scan(&key)
		switch key {
		case 1:
			fmt.Println("登录")
			fmt.Println("请输入id")
			fmt.Scan(&userId)
			fmt.Println("请输入密码")
			fmt.Scan(&userPwd)
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
			//loop = false
		case 2:
			fmt.Println("注册")
			fmt.Println("请输入id")
			fmt.Scan(&userId)
			fmt.Println("请输入密码")
			fmt.Scan(&userPwd)
			fmt.Println("请输入昵称")
			fmt.Scan(&userName)
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
			//loop = false
		case 3:
			fmt.Println("退出")
			os.Exit(3)
		default:
			fmt.Println("请重新输入")
		}
	}
	/*	if key == 1 {
			fmt.Println("请输入id")
			fmt.Scan(&id)
			fmt.Println("请输入密码")
			fmt.Scan(&password)
			login(id, password)
		} else if key == 2 {

		}*/
}
