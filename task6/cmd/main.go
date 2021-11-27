// Сделать кастомную waitGroup на семафоре

package main

import (
	"fmt"
)

type semaphore chan struct{}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := len(numbers)

	sem := New(n)

	for _, num := range numbers {
		go func(n int) {
			fmt.Println(n)
			sem.Inc(1)
		}(num)
	}

	sem.Dec(n)

}

func New(n int) semaphore {
	return make(semaphore, n)
}

func (s semaphore) Inc(k int) {
	for i := 0; i < k; i++ {
		s <- struct{}{}
	}
}

func (s semaphore) Dec(k int) {
	for i := 0; i < k; i++ {
		<-s
	}
}
