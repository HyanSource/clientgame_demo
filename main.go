package main

import (
	"fmt"

	"github.com/hyansource/20200212gameclient/data"
)

func main() {

	client := data.NewTcpClient("127.0.0.1", 7777)
	client.Start()

	fmt.Println("server succees")

	select {}
}
