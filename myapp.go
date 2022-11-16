package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	countOfFlows := flag.Int("flow", 3, "help message for flow")
	maxNumber := flag.Int("nums", 15, "help message for nums")
	flag.Parse()
	numberChan := make(chan int, *countOfFlows)
	stop := make(chan bool)
	for i := 0; i < *countOfFlows; i++ {
		go random(maxNumber, numberChan, stop, wg)
	}
	wg.Add(1)
	printer(numberChan, maxNumber, stop, wg)
	wg.Wait()
}

// Random генерирует случайные числа в диапазоне от -x до x в бесконечном цикле, и передаёт значения в буферизированный канал
func random(maxNumber *int, numberChan chan int, stop chan bool, wg *sync.WaitGroup)  {
	for {
		select {
		case <-stop:
			return
		default:
			min := -*maxNumber
			max := *maxNumber
			random := rand.Intn(max-min) + min
			numberChan <- random
		}
	}
}


/* Printer получает случайные числа из random, далее заполняет их в срез. Закрывает канал stop (завершает работу горутин), если чисел больше чем maxNumber */
func printer(numberChan chan int, maxNumber *int, stop chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	numbers := make(map[int]bool)

	count := 0

	for {
				randomNumber := <- numberChan
				if numbers[randomNumber] {
					continue
				}
				numbers[randomNumber] = true
				count++

		if count == *maxNumber {
			close(stop)
			break
		}
	

	}
	for number, _ := range numbers {
    fmt.Println(number)
  }
}
