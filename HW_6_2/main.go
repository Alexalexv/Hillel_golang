package main

import (
	"fmt"
	"sync"
)

func Publisher(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)

}

func SliceOfChanelsInit(n int) []chan int {
	sl := make([]chan int, n)
	for i := range sl {
		sl[i] = make(chan int)
	}
	return sl
}

func ChanelBroadkaster(ch <-chan int, chSl []chan int, n int) {
	for i := range ch {
		//передаємо всі повідомлення в кожен канал слайсу
		for a := 0; a < n; a++ {
			chSl[a] <- i
		}
	}
	//закриваємо кожен канал після передачі
	for a := 0; a < n; a++ {
		close(chSl[a])
	}

}

func Sucscriber(chSl []chan int, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range chSl[n] {
		fmt.Printf("Receiver #%d: %d\n", n, i)
	}
}

func main() {
	var (
		wg sync.WaitGroup
		n  int = 3
	)
	chanelMain := make(chan int)
	chanelsSlice := SliceOfChanelsInit(n)

	go ChanelBroadkaster(chanelMain, chanelsSlice, n)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go Sucscriber(chanelsSlice, i, &wg)
	}

	wg.Add(1)
	go Publisher(chanelMain, &wg)
	wg.Wait()

}
