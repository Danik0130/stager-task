package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var Port = ":5555"

func main() {

	http.HandleFunc("/", ServeFiles)
	fmt.Println("Serving @ : ", "http://127.0.0.1"+Port)
	log.Fatal(http.ListenAndServe(Port, nil))
}

func ServeFiles(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":

		path := r.URL.Path

		fmt.Println(path)

		if path == "/" {

			path = "./static/index.html"
		} else {

			path = "." + path
		}

		http.ServeFile(w, r, path)

	case "POST":

		r.ParseMultipartForm(0)

		flows, _ := strconv.Atoi(r.FormValue("flows"))
		maxNumber, _ := strconv.Atoi(r.FormValue("maxNumber"))

		fmt.Println("----------------------------------")
		fmt.Println("Messages from Client: ", flows, maxNumber)
		// respond to client's request
		fmt.Fprintf(w, "Server: %s \n", strconv.Itoa(flows)+" | "+time.Now().Format(time.RFC3339))
		Generator(flows, maxNumber)

	default:
		fmt.Fprintf(w, "Request type other than GET or POSt not supported")

	}

}
