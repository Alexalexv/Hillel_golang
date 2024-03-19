package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Publisher(ch chan<- int, wg *sync.WaitGroup) {

	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- rand.Int()
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(990)+10))
	}
}

func Subscriber(ch <-chan int) {
	for data := range ch {
		fmt.Println("Recived: ", data)
	}

}

func main() {
	var wg sync.WaitGroup

	chInt := make(chan int)

	wg.Add(1)

	go Publisher(chInt, &wg)
	go Subscriber(chInt)
	wg.Wait()
}
