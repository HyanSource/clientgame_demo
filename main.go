package main

import (
	"fmt"
	"time"

	"github.com/hyansource/20200212gameclient/data"
)

func main() {

	for i := 0; i < 300; i++ {
		client := data.NewTcpClient("127.0.0.1", 7777)
		client.Start()

		time.Sleep(time.Second / 10)
	}

	fmt.Println("server succees")

	select {}
}
