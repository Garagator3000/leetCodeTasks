//Слить N каналов в один

package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		chanA = make(chan int)
		chanB = make(chan int)
		chanC = make(chan int)
	)

	go func() {
		for i := 0; i < 5; i++ {
			chanA <- i
		}
		close(chanA)
	}()
	go func() {
		for i := 5; i < 10; i++ {
			chanB <- i
		}
		close(chanB)
	}()
	go func() {
		for i := 10; i < 15; i++ {
			chanC <- i
		}
		close(chanC)
	}()


	for num := range JoinChan(chanA, chanB, chanC) {
		fmt.Println(num)
	}
}

func JoinChan(chans ...<-chan int) <-chan int {
	merged := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(len(chans))

		for _, ch := range chans {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for i := range ch {
					merged <- i
				}
			}(ch, wg)
		}
		wg.Wait()
		close(merged)
	}()
	return merged
}
