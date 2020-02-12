package main

import (
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

func OnConnectionAdd(conn ziface.IConnection) {

	fmt.Println("client succees")
}

func OnConnectionLost(conn ziface.IConnection) {

	fmt.Println("client lost")
}

func main() {

	//創建服務器句柄
	ser := znet.NewServer()

	//註冊客戶端連接和丟失函數
	ser.SetOnConnStart(OnConnectionAdd)
	ser.SetOnConnStop(OnConnectionLost)

	//註冊路由
	//ser.AddRouter(2)
	//ser.AddRouter(3)

	s.Serve()

	fmt.Println("")
}
