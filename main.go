package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var fileWay = flag.String("File", "links.txt", "help message for File")  //flag для файла который парсим
	var fileDownloadWay = flag.String("dst", "html", "help message fot dst") //flag для директории куда скачиваем
	flag.Parse()
	file, err := os.Open(*fileWay) //открываем файл по директории из терминала
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var i int64 = 1
	scanner := bufio.NewScanner(file) //построчно считываем файл
	for scanner.Scan() {
		url := scanner.Text()
		fmt.Println(url)
		filename := *fileDownloadWay + strconv.FormatInt(i, 10) + ".html"
		download(url, filename)
		i += 1
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

func download(url, filename string) (err error) { // функция скачивания
	fmt.Println("Downloading ", url, " to ", filename)

	resp, err := http.Get(url) //создаём запрос HTTP GET
	if err != nil {
		return
	}
	defer resp.Body.Close()

	f, err := os.Create(filename) // создаём файл в директории из терминала
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body) // копируем содержимое из тела html структуры в созданный файл
	return
}
