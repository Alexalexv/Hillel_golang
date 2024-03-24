package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Publisher(ch chan<- int, wg *sync.WaitGroup, goRutineNumder int, guard chan struct{}) {
	defer wg.Done()
	fmt.Printf("gorutine #%d is started\n", goRutineNumder)
	for i := 0; i < 10; i++ {
		ch <- rand.Int()
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(990)+10))
	}
	fmt.Printf("gorutine #%d is finished\n", goRutineNumder)
	<-guard

}

func Subscriber(ch <-chan int) {
	for data := range ch {
		fmt.Println("Recived: ", data)
	}

}

func main() {
	var wg sync.WaitGroup
	guard := make(chan struct{}, 10)
	chInt := make(chan int)

	go Subscriber(chInt)

	for i := 0; i < 100; i++ {
		guard <- struct{}{}
		wg.Add(1)
		go Publisher(chInt, &wg, i, guard)
	}

	wg.Wait()

}
