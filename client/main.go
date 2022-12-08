package main

import (
	"fmt"
	"os"
)

var id int
var password string

func main() {
	var key int
	var loop = true
	for loop {
		fmt.Println("-------欢迎登录聊天系统-------")
		fmt.Println("--------1.登录聊天室---------")
		fmt.Println("---------2.注册用户---------")
		fmt.Println("--------3.退出聊天室---------")
		fmt.Println("请输入1-3")
		fmt.Scan(&key)
		switch key {
		case 1:
			fmt.Println("登录")
			loop = false
		case 2:
			fmt.Println("注册")
			loop = false
		case 3:
			fmt.Println("退出")
			os.Exit(3)
		default:
			fmt.Println("请重新输入")
		}
	}
	if key == 1 {
		fmt.Println("请输入id")
		fmt.Scan(&id)
		fmt.Println("请输入密码")
		fmt.Scan(&password)
		login(id, password)
	} else if key == 2 {

	}
}
