package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		// 握手的过程中，涉及到跨域问题的时候，允许跨域(例如从zhibo.com跨域到websocket.com)
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

func wxHandle(w http.ResponseWriter, r *http.Request) {
	var (
		conn *websocket.Conn
		err  error
		data []byte
	)
	// Upgrade:websocket
	if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	// websocket.Conn
	for {
		// Text,Binary
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}

	// 标签
ERR:
	conn.Close()
}

func main() {
	// 路由
	http.HandleFunc("/wx", wxHandle)
	http.ListenAndServe("0.0.0.0:7777", nil)
	fmt.Println("开始")
}
