package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	var fileWay = flag.String("File", "links.txt", "help message for File")  //flag для файла который парсим
	var fileDownloadWay = flag.String("dst", "html", "help message fot dst") //flag для директории куда скачиваем
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
	size := len(links)
	ch := make(chan int, size) // создаём каналы чтобы синхронизировать горутины
	out := make(chan int, size)
	for n, _ := range links {

		filename := *fileDownloadWay + strconv.FormatInt(i, 10) + ".html"
		ch <- n // отправляем в канал значение n
		go func() {
			download(links[<-ch] /*получаем из канала значение n */, filename)
			out <- 0
		}()
		i++
	}
	var count int
	for { // счётчик, который останавливает запуск горутин, когда кончился срез links
		select {
		case <-out:
			count++
			if count == size {
				return
			}
		default:
			time.Sleep(1 * time.Millisecond)
		}
	}
}

// download скачивает файл со структурой html в указанную пользователем директорию
func download(url string, filename string) (err error) {
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
