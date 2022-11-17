package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var Port = ":5555" // установили порт сервера

func main() {
 
	http.HandleFunc("/", ServeFiles) // создали роутер
	fmt.Println("Server: ", "http://127.0.0.1"+Port)
	log.Fatal(http.ListenAndServe(Port, nil))
}

func ServeFiles(w http.ResponseWriter, r *http.Request) {

	switch r.Method {  // переключатель, который считывает формат запроса

	case "GET": // при запросе GET открываем сайт по абсолютному адресу (для задания по сути не нужно)

		path := r.URL.Path

		fmt.Println(path)

		if path == "/" {

			path = "./static/index.html"
		} else {

			path = "." + path
		}

		http.ServeFile(w, r, path)

	case "POST": // при запросе POST получаем данные из формы, выполняем функцию Generator

		r.ParseMultipartForm(0) 

		flows, _ := strconv.Atoi(r.FormValue("flows")) // получаем значение потоков с формы
		maxNumber, _ := strconv.Atoi(r.FormValue("maxNumber")) // получаем максимальное число с формы

		fmt.Println("----------------------------------")
		fmt.Println("Messages from Client: ", flows, maxNumber)
		// ответ на запрос клиента
		fmt.Fprintf(w, "Server: %s \n", strconv.Itoa(flows)+" "+strconv.Itoa(maxNumber))
		Generator(flows, maxNumber)

	default:
		fmt.Fprintf(w, "Request type other than GET or POSt not supported")

	}

}
