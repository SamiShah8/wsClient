package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for i := 0; i < 10; i++ {
		go connectWs()
	}

	<-interrupt

}

func connectWs() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://192.168.1.100:9090/ws", nil)
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		conn.Close()
	}()

	for {
		err := conn.WriteMessage(websocket.BinaryMessage, []byte(fmt.Sprintf("count: %d", 1)))
		if err != nil {
			fmt.Println(err)
			return
		}

		_, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("msg: %v\n", string(data))
		time.Sleep(100 * time.Millisecond)
	}
}
