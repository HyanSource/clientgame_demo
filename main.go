package main

import (
	"fmt"

	"github.com/hyansource/clientgame_demo/data"
)

func main() {

	//for i := 0; i < 300; i++ {
	client := data.NewTcpClient("127.0.0.1", 7777)
	client.Start()

	//time.Sleep(time.Second / 10)
	//}

	fmt.Println("server succees")

	select {}
}
