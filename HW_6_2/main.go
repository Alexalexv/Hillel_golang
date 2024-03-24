package main

import (
	"fmt"
	"sync"
)

func Publisher(ch chan<- int, wg *sync.WaitGroup) {
	defer close(ch)
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i * 2
	}
}

func ChanelBroadkaster(ch <-chan int, chSl []chan int, n int) {
	for i := range ch {
		for a := 0; a < n; a++ {
			val := i // Створюємо копію значення i
			chSl[a] <- val
		}
	}
}

func Sucscriber(ch <-chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Printf("Receiver #%d: %d\n", id, i)
	}
}

func main() {
	var wg sync.WaitGroup
	n := 10
	chanelMain := make(chan int)
	chanelsSlice := make([]chan int, n)

	go ChanelBroadkaster(chanelMain, chanelsSlice, n)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go Sucscriber(chanelsSlice[i], i, &wg)
	}

	wg.Add(1)
	go Publisher(chanelMain, &wg)

	wg.Wait()
}
