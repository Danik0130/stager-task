package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	countOfFlows := flag.Int("flow", 3, "help message for flow")
	maxNumber := flag.Int("nums", 15, "help messsage for nums")
	flag.Parse()
	numberChan := make(chan int, *countOfFlows)
	min := -*maxNumber
	max := *maxNumber
	for i := 0; i < *countOfFlows; i++ {
		wg.Add(1)
		go func() {
			random := rand.Intn(max-min) + min
			numberChan <- random
		}()

		func ()  {
		defer wg.Done()
		fmt.Println(<-numberChan)
	}()
	}

wg.Wait()
}
