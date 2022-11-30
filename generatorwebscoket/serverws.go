package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Пропускаем любой запрос
	},
}
// gen обрабатывает сокет соединение
func gen(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade connection from %s error", r.RemoteAddr)
		return
	}
	go listener(connection)
}

// listener получает кол-во потоков и максимальное число, затем передаёт их в Generator
func listener(connection *websocket.Conn) {
	defer connection.Close()

	type num struct {
		Flows     int `json:"flows"`
		MaxNumber int `json:"maxNumber"`
	}
	type res struct {
		Result []int  `json:"result"`
	}
	for {
		var message num
		err := connection.ReadJSON(&message)

		if err != nil {
			log.Println(err)
			break // Если ошибка обрываем соединение
		}
		var response res
		genOutput := Generator(message.Flows, message.MaxNumber)
		for number, _ := range genOutput {
			response.Result = append(response.Result, number)
		}

		err = connection.WriteJSON(response)
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fs := http.FileServer(http.Dir("./open"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/generator", gen)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "open/index.html")
	})


	log.Println("Connection established")


	err := http.ListenAndServe(":5555", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

