package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/", WebSocket)
	http.ListenAndServe(":5555", nil)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Пропускаем любой запрос
	},
}

func WebSocket(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer connection.Close()
	for {
		mt, flows, err := connection.ReadMessage()
		mt, maxNumber, err := connection.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			break // Выходим из цикла, если клиент пытается закрыть соединение или связь с клиентом прервана
		}

		// connection.WriteMessage(websocket.TextMessage, flows)
		a, err := strconv.Atoi(string(flows))
		b, err := strconv.Atoi(string(maxNumber))
		go Generator(a, b)
	}
}
