package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func main() {
	var fileWay = flag.String("File", "links.txt", "help message for File")  //flag для файла который парсим
	var fileDownloadWay = flag.String("dst", "html", "help message fot dst") //flag для директории куда скачиваем
	wg :=  new(sync.WaitGroup)
	flag.Parse()
	file, err := os.Open(*fileWay)
	if err != nil {
		os.Exit(1)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var links []string

	for scanner.Scan() {
		links = append(links, scanner.Text())
	}
	fmt.Println(links)
	if err != nil {
		os.Exit(1)
		return
	}
	var i int64 = 1
	for _, url := range links {
		wg.Add(1)
		filename := *fileDownloadWay + strconv.FormatInt(i, 10) + ".html"
		go download(url, filename, wg)

		i++
	}
	wg.Wait()
}

// download скачивает файл со структурой html в указанную пользователем директорию
func download(url string, filename string, wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	fmt.Println("Downloading ", url, " to ", filename)
	resp, err := http.Get(url) // создаём запрос HTTP GET
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
	return nil
}
