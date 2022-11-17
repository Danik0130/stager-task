package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func Generator(flows int, maxNumber int) {
	wg := new(sync.WaitGroup)
	if flows <= 0 || flows > 100 {
		fmt.Println("Error: wrong number of flows")
		return
	}
	if maxNumber <= 0 || maxNumber > 100000 {
		fmt.Println("Wrong Max Number")
		return
	}
	numberChan := make(chan int, flows)
	stop := make(chan bool)
	for i := 0; i < flows; i++ {
		go random(maxNumber, numberChan, stop)
	}
	wg.Add(1)
	printer(numberChan, maxNumber, stop, wg)
	wg.Wait()
}

// Random генерирует случайные числа в диапазоне от 0 до x в бесконечном цикле, и передаёт значения в буферизированный канал
func random(maxNumber int, numberChan chan int, stop chan bool) {
	for {
		select { // селектор дожидается отправки данных (закрытия канала), после чего горутины остановятся
		case <-stop:
			return
		default: // если канал не закрыт горутины продолжают работать
			max := maxNumber
			random := rand.Intn(max)
			numberChan <- random
		}
	}
}

/* Printer получает случайные числа из random, далее заполняет их в срез. Закрывает канал stop (завершает работу горутин), если чисел больше чем maxNumber */
func printer(numberChan chan int, maxNumber int, stop chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	numbers := make(map[int]bool) // создаём map для проверки того, что число уже попадалось

	count := 0

	for {
		randomNumber := <-numberChan
		if numbers[randomNumber] { // приступаем к следующей итерации если число уже было
			continue
		}
		numbers[randomNumber] = true
		count++

		if count == maxNumber { /* если количество чисел = *maxNumber, то закрываем канал (завершаем гортины) и завершаем цикл */
			close(stop)
			break
		}

	}
	for number, _ := range numbers {
		fmt.Println(number)
	}
}
